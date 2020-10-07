package zerotears

import (
	"encoding/json"
	"github.com/rodaine/table"
)

type networkInfo struct {
	CreationTime      int                `json:"creationTime,omitempty"`
	EnableBroadcast   bool               `json:"enableBroadcast,omitempty"`
	ID                string             `json:"id,omitempty"`
	IPAssignmentPools []ipAssignmentPool `json:"ipAssignmentPools,omitempty"`
	Name              string             `json:"name,omitempty"`
	Private           bool               `json:"private,omitempty"`
	MTU               int                `json:"mtu,omitempty"`
	MulticastLimit    int                `json:"multicastLimit,omitempty"`
	Revision          int                `json:"revision,omitempty"`
	Routes            []route            `json:"routes,omitempty"`
	V4AssignMode      *v4AssignMode      `json:"v4AssignMode,omitempty"`
}

type v4AssignMode struct {
	Zt bool `json:"zt,omitempty"`
}

type ipAssignmentPool struct {
	IPRangeEnd   string `json:"ipRangeEnd,omitempty"`
	IPRangeStart string `json:"ipRangeStart,omitempty"`
}

type route struct {
	Target string `json:"target,omitempty"`
	Via    string `json:"via,omitempty"`
}

func (z *ztClient) DeleteNetwork(id string) networkInfo {
	var out networkInfo
	z.doReq("DELETE", "/controller/network/"+id, nil, &out)
	return out
}

func (z *ztClient) CreateNetwork(name, rangeStart, rangeEnd, cidr string, autoAssign bool) networkInfo {
	newNet := networkInfo{
		Name: name,
		V4AssignMode: &v4AssignMode{
			Zt: autoAssign,
		},
		IPAssignmentPools: []ipAssignmentPool{
			{
				IPRangeStart: rangeStart,
				IPRangeEnd:   rangeEnd,
			},
		},
		Routes: []route{
			{
				Target: cidr,
			},
		},
	}
	payload, err := json.Marshal(newNet)
	if err != nil {
		panic(err)
	}

	var out networkInfo
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

func (z *ztClient) ListNetworksVerbose() []networkInfo {
	var out []networkInfo
	for _, v := range z.ListNetworks() {
		out = append(out, z.GetNetworkInfo(v))
	}
	return out
}

func PrintNetworkInfo(in []networkInfo) {
	tbl := table.New("Network ID", "Name", "Private", "EnableBroadcast", "IPv4 Auto Assign", "IP Range Start", "IP Range End")
	for _, v := range in {
		tbl.AddRow(v.ID, v.Name, v.Private, v.EnableBroadcast, v.V4AssignMode.Zt, v.IPAssignmentPools[0].IPRangeStart, v.IPAssignmentPools[0].IPRangeEnd)
	}
	tbl.Print()
}

func PrintNetworkIDs(in []string) {
	tbl := table.New("Network ID")
	for _, v := range in {
		tbl.AddRow(v)
	}
	tbl.Print()
}
