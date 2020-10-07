package zerotears

import (
	"github.com/rodaine/table"
)

type controllerInfoResp struct {
	Controller    bool
	ApiVersion    int
	Clock         int
	DatabaseReady bool
}

type statusResp struct {
	Address              string
	Clock                int
	Online               bool
	PlanetWorldID        int
	PlanetWorldTimestamp int
	PublicIdentity       string
	TCPFallbackActive    bool
	Version              string
}

func (z *ztClient) GetControllerInfo() controllerInfoResp {
	var out controllerInfoResp
	z.doReq("GET", "/controller", nil, &out)
	return out
}

func (z *ztClient) GetStatus() statusResp {
	var out statusResp
	z.doReq("GET", "/status", nil, &out)
	return out
}

func PrintControllerInfo(in *controllerInfoResp) {
	tbl := table.New("Controller", "ApiVersion", "Clock", "DatabaseReady")
	tbl.AddRow(in.Controller, in.ApiVersion, in.Clock, in.DatabaseReady)
	tbl.Print()
}

func PrintStatus(in *statusResp) {
	tbl := table.New("Address", "Online", "PlanetWorldID", "TCPFallbackActive", "Version")
	tbl.AddRow(in.Address, in.Online, in.PlanetWorldID, in.TCPFallbackActive, in.Version)
	tbl.Print()
}
