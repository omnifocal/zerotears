package main

import (
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

func (z *ztClient) deleteNetwork(id string) networkInfo {
	var out networkInfo
	z.doReq("DELETE", "/controller/network/"+id, nil, &out)
	return out
}

func (z *ztClient) createNetwork(name string) networkInfo {
	var out networkInfo
	payload := []byte(`{"name":"` + name + `"}`)
	z.doReq("POST", "/controller/network/"+z.address+"______", payload, &out)
	return out
}

func (z *ztClient) getNetworkInfo(id string) networkInfo {
	var out networkInfo
	z.doReq("GET", "/controller/network/"+id, nil, &out)
	return out
}

func (z *ztClient) listNetworks() []string {
	var out []string
	z.doReq("GET", "/controller/network", nil, &out)
	return out
}

func printNetworkIDs(in []string) {
	tbl := table.New("Network ID")
	for _, v := range in {
		tbl.AddRow(v)
	}
	tbl.Print()
}
