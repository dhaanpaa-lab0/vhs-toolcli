package host

import (
	"github.com/go-ini/ini"
	"log"
	"nexus-csd.net/vhs-toolcli/utils"
	"os"
	"os/user"
)

func GetWebServerUser() (string, error) {
	usr, err := user.Lookup(GetWebServerUserGroup())
	if err != nil {
		log.Fatal("error: Retrieving User: ", err)
		return "", err

	} else {
		return usr.Uid, nil
	}
	return "", nil
}

func GetWebServerGroup() (string, error) {
	grp, err := user.LookupGroup(GetWebServerUserGroup())
	if err != nil {
		log.Fatal("error: Retrieving Group: ", err)
		return "", err

	} else {
		return grp.Gid, nil
	}
	return "", nil
}

func GetWebServerUserGroup() string {
	if IsUbuntu() {
		return "www-data"
	}

	if IsRedhatBased() {
		return "apache"
	}

	return ""
}

func readOSRelease(configfile string) map[string]string {
	ConfigParams := make(map[string]string)
	if !utils.Exists(configfile) {
		ConfigParams["ID"] = ""
		return ConfigParams
	}

	cfg, err := ini.Load(configfile)

	if err != nil {
		ConfigParams["ID"] = ""
		log.Println("Failed to read file: ", err)
		return ConfigParams
	}

	ConfigParams["ID"] = cfg.Section("").Key("ID").String()

	return ConfigParams
}

func GetSitesAvailableFolder() string {
	amazonSitesAvailableFolder := "/etc/httpd/sites-available"
	ubuntuSitesAvailableFolder := "/etc/apache2/sites-available"
	if IsRedhatBased() {
		if !utils.Exists(amazonSitesAvailableFolder) {
			err := os.MkdirAll(amazonSitesAvailableFolder, 0755)
			if err != nil {
				log.Println(err)
				return ""
			} else {
				return amazonSitesAvailableFolder
			}
		} else {
			return amazonSitesAvailableFolder
		}
	}

	if IsUbuntu() {
		if !utils.Exists(ubuntuSitesAvailableFolder) {
			err := os.MkdirAll(ubuntuSitesAvailableFolder, 0755)
			if err != nil {
				log.Println(err)
				return ""
			} else {
				return ubuntuSitesAvailableFolder
			}
		} else {
			return ubuntuSitesAvailableFolder
		}
	}
	return ""
}

func IsRedhatBased() bool {
	if GetOs() == "aws" {
		return true
	} else {
		return false
	}
}

func IsUbuntu() bool {
	if GetOs() == "ubuntu" {
		return true
	} else {
		return false
	}
}

func GetOs() string {
	osInfo := readOSRelease("/etc/os-release")
	osId := osInfo["ID"]

	switch osId {
	case "amzn":
		return "aws"
	case "ubuntu":
		return "ubuntu"
	default:
		return "unknown"
	}
}

func GetSitesEnabledFolder() string {
	amazonSitesEnabledFolder := "/etc/httpd/conf.d"
	ubuntuSitesEnabledFolder := "/etc/apache2/sites-enabled"
	if IsRedhatBased() {
		if !utils.Exists(amazonSitesEnabledFolder) {
			err := os.MkdirAll(amazonSitesEnabledFolder, 0755)
			if err != nil {
				log.Println(err)
				return ""
			} else {
				return amazonSitesEnabledFolder
			}
		} else {
			return amazonSitesEnabledFolder
		}
	}

	if IsUbuntu() {
		if !utils.Exists(ubuntuSitesEnabledFolder) {
			err := os.MkdirAll(ubuntuSitesEnabledFolder, 0755)
			if err != nil {
				log.Println(err)
				return ""
			} else {
				return ubuntuSitesEnabledFolder
			}
		} else {
			return ubuntuSitesEnabledFolder
		}
	}
	return ""
}
