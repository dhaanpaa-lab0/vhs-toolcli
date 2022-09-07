package main

import (
	"flag"
	"fmt"
	"nexus-csd.net/vhs-toolcli/commands"
	"nexus-csd.net/vhs-toolcli/host"
	"nexus-csd.net/vhs-toolcli/server"
	"os"
)

func main() {

	cmd := flag.String("cmd", "", "Enter Action (create,enable,disable,delete)")
	domain := flag.String("domain", "", "Enter the virtual host domain name.(Required)")
	flag.Parse()
	fmt.Println("Operating System .......... ", host.GetOs())
	fmt.Println("Sites Enabled Folder ...... ", host.GetSitesEnabledFolder())
	fmt.Println("Sites Available Folder .... ", host.GetSitesAvailableFolder())

	if host.GetOs() == "unknown" {
		fmt.Println("Unsupported OS")
		os.Exit(1)
	}

	if *domain == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *cmd == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	commandToExecute := commands.ToVhsCommand(*cmd)
	switch commandToExecute {
	case commands.CreateSite:
		server.CreateVirtualSite(*domain)
	case commands.DeleteSite:
		server.DeleteVirtualSite(*domain)
	case commands.EnableSite:
		server.EnableVirtualSite(*domain)
	case commands.DisableSite:
		server.DisableVirtualSite(*domain)
	default:
		fmt.Println("Unkonw Command: " + commandToExecute)
	}
}
