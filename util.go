package main

import (
	"io/ioutil"
	"net/http"
)

type ztClient struct {
	host    string
	secret  string
	address string
	client  http.Client
}

func (z *ztClient) init() {
	status := z.getStatus()
	z.address = status.Address
}

func (z *ztClient) doReq(req *http.Request) *http.Response {
	req.Header.Add("X-ZT1-Auth", z.secret)

	resp, err := z.client.Do(req)
	if err != nil {
		panic(err)
	}
	return resp
}

func readSecret(fileName string) string {
	dat, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(dat)
}
