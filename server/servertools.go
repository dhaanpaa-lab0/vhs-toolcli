package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"nexus-csd.net/vhs-toolcli/host"
	"os"
	"path/filepath"
	"strconv"
)

func CreateServerFolder(domain string) {
	serverUser, createFolder := host.GetWebServerUser()
	if createFolder != nil {
		return
	}

	serverGroup, createFolder := host.GetWebServerGroup()
	if createFolder != nil {
		return
	}

	var createFolderErr = os.MkdirAll(GetWebHostFilesPath(domain), 0755)

	if createFolderErr != nil {
		fmt.Println(createFolderErr)
		os.Exit(1)
	} else {
		fmt.Println("Directory and permissions for the " + domain + " created successfully. \n")
	}

	Uid, createFolderErr := strconv.ParseInt(serverUser, 10, 64)
	userId := int(Uid)

	Gid, createFolderErr := strconv.ParseInt(serverGroup, 10, 64)
	groupId := int(Gid)

	fmt.Printf("Setting file ownership for the new domain. \n")
	createFolderErr = os.Chown(GetWebHostFilesPath(domain), userId, groupId)
}

func GetWebHostFilesPath(domain string) string {
	return "/var/www/" + domain + "/public_html"
}

func GetHttpLogFolder() string {
	if host.IsUbuntu() {
		return "/var/log/apache2"
	}

	if host.IsRedhatBased() {
		return "/var/log/httpd"
	}

	return ""
}
func CreateHostConfig(domain string) {
	virtualHostSettings := `<VirtualHost *:80>` + "\n\n" +
		`  ServerAdmin admin@` + domain + "\n" +
		`  ServerName ` + domain + "\n" +
		`  DocumentRoot /var/www/` + domain + `/public_html` + "\n\n" +
		`  <Directory /var/www/` + domain + `/public_html>` + "\n" +
		`     Options -Indexes +FollowSymLinks -MultiViews` + "\n" +
		`     AllowOverride All ` + "\n" +
		`     Require all granted` + "\n" +
		`   </Directory>` + "\n\n" +
		`   ErrorLog ` + GetHttpLogFolder() + `/` + domain + `.error.log` + "\n" +
		`   CustomLog ` + GetHttpLogFolder() + `/` + domain + `.access.log combined` + "\n\n" +
		` </VirtualHost>`

	virtualHostConfigBytes := []byte(virtualHostSettings)

	fmt.Printf("Creating virtual host configuration file... \n")

	err := ioutil.WriteFile(GetSitesAvailableConfigFile(domain), virtualHostConfigBytes, 0644)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Printf("Virtual host configuration file created. \n")
	}
}

func GetSitesAvailableConfigFile(domain string) string {
	return filepath.Join(host.GetSitesAvailableFolder(), GetConfigFileName(domain))
}

func GetSitesEnabledConfigFile(domain string) string {
	return filepath.Join(host.GetSitesEnabledFolder(), GetConfigFileName(domain))
}

func GetConfigFileName(domain string) string {
	return domain + ".conf"
}

func CreateVirtualSite(domain string) {
	fmt.Println("Domain ............. " + domain)
	CreateServerFolder(domain)
	CreateHostConfig(domain)
	EnableVirtualSite(domain)
	RestartServer()
}

func RestartServer() {
}

func DeleteVirtualServerFolder(domain string) {
	err := os.RemoveAll(GetWebHostFilesPath(domain))
	if err != nil {
		log.Println(err)
		return
	} else {
		log.Println("Site " + domain + " Folder Removed Successfully")
	}
}

func DeleteVirtualServerConfig(domain string) {
	err := os.Remove(GetSitesAvailableConfigFile(domain))
	if err != nil {
		log.Println(err)
		return
	} else {
		log.Println("Site " + domain + " Config Removed Successfully")
	}
}

func DeleteVirtualSite(domain string) {
	DisableVirtualSite(domain)
	DeleteVirtualServerFolder(domain)
	DeleteVirtualServerConfig(domain)
	RestartServer()
}

func EnableVirtualSite(domain string) {
	err := os.Symlink(GetSitesAvailableConfigFile(domain), GetSitesEnabledConfigFile(domain))
	if err != nil {
		log.Println(err)
		return
	} else {
		log.Println("Site " + domain + " Enabled Successfully")
	}
	RestartServer()
}

func DisableVirtualSite(domain string) {
	err := os.Remove(GetSitesEnabledConfigFile(domain))
	if err != nil {
		log.Println(err)
		return
	} else {
		log.Println("Site " + domain + "Disabled Successfully")
	}
	RestartServer()
}
