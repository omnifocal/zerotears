package main

import (
	"encoding/json"
	"fmt"
	"github.com/rodaine/table"
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
	resp := z.doReq("DELETE", "/controller/network/"+id, nil)
	fmt.Println(string(resp))
}

func (z *ztClient) createNetwork(name string) {
	payload := []byte(`{"name":"` + name + `"}`)
	resp := z.doReq("POST", "/controller/network/"+z.address+"______", payload)
	fmt.Println(string(resp))
}

func (z *ztClient) getNetworkInfo(id string) networkInfo {
	resp := z.doReq("GET", "/controller/network/"+id, nil)

	var out networkInfo
	err := json.Unmarshal(resp, &out)
	if err != nil {
		panic(err)
	}
	return out
}

func (z *ztClient) listNetworks() []string {
	resp := z.doReq("GET", "/controller/network", nil)

	var out []string
	err := json.Unmarshal(resp, &out)
	if err != nil {
		panic(err)
	}
	return out
}

func printNetworkIDs(in []string) {
	tbl := table.New("Network ID")
	for _, v := range in {
		tbl.AddRow(v)
	}
	tbl.Print()
}
