package main

import (
	"flag"
	"fmt"
	"github.com/omnifocal/zerotears/pkg/zerotears"
	"io/ioutil"
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
	// listNetsCmd := flag.NewFlagSet("ls-networks", flag.ExitOnError)
	createNetCmd := flag.NewFlagSet("mk-network", flag.ExitOnError)
	createNetName := createNetCmd.String("name", "", "name")

	switch os.Args[1] {
	case "status":
		status := c.GetStatus()
		zerotears.PrintStatus(&status)
	case "info":
		info := c.GetControllerInfo()
		zerotears.PrintControllerInfo(&info)
	case "ls-networks":
		networks := c.ListNetworks()
		zerotears.PrintNetworkIDs(networks)
	case "mk-network":
		createNetCmd.Parse(os.Args[2:])
		newNet := c.CreateNetwork(*createNetName)
		fmt.Println(newNet)
	default:
		fmt.Println("No valid subcommand was provided")
		os.Exit(1)
	}
}
