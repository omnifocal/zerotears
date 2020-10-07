package main

import (
	"github.com/omnifocal/zerotears/pkg/zerotears"
	"io/ioutil"
)

func main() {
	host := "http://127.0.0.1:9993"
	dat, err := ioutil.ReadFile("authtoken.secret")
	if err != nil {
		panic(err)
	}

	c := zerotears.Init(host, string(dat))

	info := c.GetStatus()
	zerotears.PrintStatus(&info)

	networks := c.ListNetworks()
	zerotears.PrintNetworkIDs(networks)
}
