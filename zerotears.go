package main

import (
	"net/http"
)

func main() {
	host := "http://127.0.0.1:9993"
	secret := readSecret("./authtoken.secret")
	c := ztClient{
		host:   host,
		secret: secret,
		client: http.Client{},
	}
	c.init()

	info := c.getStatus()
	printStatus(&info)

	networks := c.listNetworks()
	printNetworkIDs(networks)
}
