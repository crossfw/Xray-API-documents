package examples

import (
	"fmt"
	"testing"
)

func TestGetSysStats(t *testing.T) {
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
	SysStats, err := getSysStats(xrayCtl.SsClient)
	if err != nil {
		t.Errorf("Failed %s", err)
	}
	fmt.Println(SysStats)

}
