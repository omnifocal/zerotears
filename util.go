package main

import (
	"bytes"
	"encoding/json"
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

func (z *ztClient) doReq(method string, path string, body []byte, out interface{}) {
	req, err := http.NewRequest(method, z.host+path, bytes.NewBuffer(body))
	req.Header.Add("X-ZT1-Auth", z.secret)
	if err != nil {
		panic(err)
	}

	resp, err := z.client.Do(req)
	if err != nil {
		panic(err)
	}

	outBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(outBytes, out)
	if err != nil {
		panic(err)
	}
}

func readSecret(fileName string) string {
	dat, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(dat)
}
