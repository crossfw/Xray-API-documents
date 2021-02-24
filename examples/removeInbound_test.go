package examples

import "testing"

func TestRemoveInbound(t *testing.T) {
	var (
		xrayCtl *XrayController
		cfg     = &BaseConfig{
			APIAddress: "127.0.0.1",
			APIPort:    10085,
		}
		tag = "proxy0"
	)
	xrayCtl = new(XrayController)
	err := xrayCtl.Init(cfg)
	defer xrayCtl.CmdConn.Close()
	if err != nil {
		t.Errorf("Failed %s", err)
	}
	err = removeInbound(xrayCtl.HsClient, tag)
	if err != nil {
		t.Errorf("Failed %s", err)
	}

}
