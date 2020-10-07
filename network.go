package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/rodaine/table"
	"io/ioutil"
	"net/http"
)

type networkInfo struct {
	CreationTime      int
	EnableBroadcast   bool
	ID                string
	IPAssignmentPools []ipAssignmentPool
	Name              string
	Private           bool
	MTU               int
	MulticastLimit    int
	Revision          int
}

type ipAssignmentPool struct {
	IPRangeEnd   string
	IPRangeStart string
}

type route struct {
	Target string
	Via    string
}

func (z *ztClient) deleteNetwork(id string) {
	url := z.host + "/controller/network/" + id
	req, _ := http.NewRequest("DELETE", url, nil)
	resp := z.doReq(req)

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func (z *ztClient) createNetwork(name string) {
	url := z.host + "/controller/network/" + z.address + "______"
	payload := []byte(`{"name":"` + name + `"}`)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	z.doReq(req)
}

func (z *ztClient) getNetworkInfo(id string) networkInfo {
	url := z.host + "/controller/network/" + id
	req, _ := http.NewRequest("GET", url, nil)
	resp := z.doReq(req)

	var out networkInfo
	body, _ := ioutil.ReadAll(resp.Body)
	err := json.Unmarshal(body, &out)
	if err != nil {
		panic(err)
	}
	return out
}

func (z *ztClient) listNetworks() []string {
	url := z.host + "/controller/network"
	req, _ := http.NewRequest("GET", url, nil)
	resp := z.doReq(req)

	out := new([]string)
	body, _ := ioutil.ReadAll(resp.Body)
	err := json.Unmarshal(body, out)
	if err != nil {
		panic(err)
	}
	return *out
}

func printNetworkIDs(in []string) {
	tbl := table.New("Network ID")
	for _, v := range in {
		tbl.AddRow(v)
	}
	tbl.Print()
}
