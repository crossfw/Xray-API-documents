package examples

import (
	"fmt"
	"testing"
)

func TestQueryTraffic(t *testing.T) {
	var (
		xrayCtl *XrayController
		cfg     = &BaseConfig{
			APIAddress: "127.0.0.1",
			APIPort:    10085,
		}
	)
	xrayCtl = new(XrayController)
	err := xrayCtl.Init(cfg)
	defer xrayCtl.CmdConn.Close()
	if err != nil {
		t.Errorf("Failed %s", err)
	}
	ptn := "inbound>>>proxy0>>>traffic>>>downlink"
	trafficData, err := queryTraffic(xrayCtl.SsClient, ptn, false)
	if err != nil {
		t.Errorf("Failed %s", err)
	}
	fmt.Println(trafficData)

}
