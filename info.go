package main

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

func (z *ztClient) getControllerInfo() controllerInfoResp {
	var out controllerInfoResp
	z.doReq("GET", "/controller", nil, &out)
	return out
}

func (z *ztClient) getStatus() statusResp {
	var out statusResp
	z.doReq("GET", "/status", nil, &out)
	return out
}

func printControllerInfo(in *controllerInfoResp) {
	tbl := table.New("Controller", "ApiVersion", "Clock", "DatabaseReady")
	tbl.AddRow(in.Controller, in.ApiVersion, in.Clock, in.DatabaseReady)
	tbl.Print()
}

func printStatus(in *statusResp) {
	tbl := table.New("Address", "Online", "PlanetWorldID", "TCPFallbackActive", "Version")
	tbl.AddRow(in.Address, in.Online, in.PlanetWorldID, in.TCPFallbackActive, in.Version)
	tbl.Print()
}
