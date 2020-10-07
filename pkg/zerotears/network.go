package zerotears

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
	Routes            []route
	V4AssignMode      v4AssignMode
}

type v4AssignMode struct {
	Zt bool
}

type ipAssignmentPool struct {
	IPRangeEnd   string
	IPRangeStart string
}

type route struct {
	Target string
	Via    string
}

func (z *ztClient) DeleteNetwork(id string) networkInfo {
	var out networkInfo
	z.doReq("DELETE", "/controller/network/"+id, nil, &out)
	return out
}

func (z *ztClient) CreateNetwork(name string) networkInfo {
	var out networkInfo
	payload := []byte(`{"name":"` + name + `"}`)
	z.doReq("POST", "/controller/network/"+z.address+"______", payload, &out)
	return out
}

func (z *ztClient) GetNetworkInfo(id string) networkInfo {
	var out networkInfo
	z.doReq("GET", "/controller/network/"+id, nil, &out)
	return out
}

func (z *ztClient) ListNetworks() []string {
	var out []string
	z.doReq("GET", "/controller/network", nil, &out)
	return out
}

func PrintNetworkInfo(in networkInfo) {
	tbl := table.New("Network ID", "Name", "Private", "EnableBroadcast", "IPv4 Auto Assign")
	tbl.AddRow(in.ID, in.Name, in.Private, in.EnableBroadcast, in.V4AssignMode.Zt)
	tbl.Print()
}

func PrintNetworkIDs(in []string) {
	tbl := table.New("Network ID")
	for _, v := range in {
		tbl.AddRow(v)
	}
	tbl.Print()
}
