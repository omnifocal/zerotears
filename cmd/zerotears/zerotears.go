package main

import (
	"flag"
	"fmt"
	"github.com/apparentlymart/go-cidr/cidr"
	"github.com/omnifocal/zerotears/pkg/zerotears"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	host := "http://127.0.0.1:9993"
	dat, err := ioutil.ReadFile("authtoken.secret")
	if err != nil {
		panic(err)
	}
	c := zerotears.Init(host, string(dat))

	// getStatusCmd := flag.NewFlagSet("status", flag.ExitOnError)
	// getInfoCmd := flag.NewFlagSet("info", flag.ExitOnError)
	listNetsCmd := flag.NewFlagSet("ls-networks", flag.ExitOnError)
	listNetsVerbose := listNetsCmd.Bool("v", false, "verbose")
	createNetCmd := flag.NewFlagSet("mk-network", flag.ExitOnError)
	createNetName := createNetCmd.String("name", "", "name")
	createNetAuto := createNetCmd.Bool("auto", true, "auto v4 assignment")
	createNetCIDR := createNetCmd.String("cidr", "", "cidr")
	rmNetCmd := flag.NewFlagSet("rm-network", flag.ExitOnError)
	rmNetID := rmNetCmd.String("id", "", "id")

	switch os.Args[1] {
	case "status":
		status := c.GetStatus()
		zerotears.PrintStatus(&status)
	case "info":
		info := c.GetControllerInfo()
		zerotears.PrintControllerInfo(&info)
	case "ls-networks":
		listNetsCmd.Parse(os.Args[2:])
		if *listNetsVerbose {
			networks := c.ListNetworksVerbose()
			zerotears.PrintNetworkInfo(networks)
		} else {
			networks := c.ListNetworks()
			zerotears.PrintNetworkIDs(networks)
		}
	case "mk-network":
		createNetCmd.Parse(os.Args[2:])
		_, ipNet, err := net.ParseCIDR(*createNetCIDR)
		if err != nil {
			panic(err)
		}
		rangeStart, rangeEnd := cidr.AddressRange(ipNet)
		c.CreateNetwork(*createNetName, rangeStart.String(), rangeEnd.String(), *createNetCIDR, *createNetAuto)
		fmt.Println("Network created")
	case "rm-network":
		rmNetCmd.Parse(os.Args[2:])
		c.DeleteNetwork(*rmNetID)
		fmt.Println("Network deleted")
	default:
		fmt.Println("No valid subcommand was provided")
		os.Exit(1)
	}
}
