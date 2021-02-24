package examples

import "testing"

func TestAddInbound(t *testing.T) {
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
		t.Errorf("Failed%s", err)
	}
	err = addInbound(xrayCtl.HsClient)
	if err != nil {
		t.Errorf("Failed%s", err)
	}

}
