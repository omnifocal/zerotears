package main

import (
	"encoding/json"
	"github.com/rodaine/table"
	"io/ioutil"
	"net/http"
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
	url := z.host + "/controller"
	req, _ := http.NewRequest("GET", url, nil)
	resp := z.doReq(req)

	var out controllerInfoResp
	body, _ := ioutil.ReadAll(resp.Body)
	err := json.Unmarshal(body, &out)
	if err != nil {
		panic(err)
	}
	return out
}

func (z *ztClient) getStatus() statusResp {
	url := z.host + "/status"
	req, _ := http.NewRequest("GET", url, nil)
	resp := z.doReq(req)

	var out statusResp
	body, _ := ioutil.ReadAll(resp.Body)
	err := json.Unmarshal(body, &out)
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
