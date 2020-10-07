package main

import (
	"encoding/json"
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
	resp := z.doReq("GET", "/controller", nil)

	var out controllerInfoResp
	err := json.Unmarshal(resp, &out)
	if err != nil {
		panic(err)
	}
	return out
}

func (z *ztClient) getStatus() statusResp {
	resp := z.doReq("GET", "/status", nil)

	var out statusResp
	err := json.Unmarshal(resp, &out)
	if err != nil {
		panic(err)
	}
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
