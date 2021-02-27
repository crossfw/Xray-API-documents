package examples

import "testing"

func TestRouteRules(t *testing.T) {
	var (
		xrayCtl *XrayController
		cfg     = &BaseConfig{
			APIAddress: "127.0.0.1",
			APIPort:    10085,
		}
	)
	xrayCtl = new(XrayController)
	err := xrayCtl.Init(cfg)
	if err != nil {
		t.Errorf("Failed%s", err)
	}
	err = addRouting(xrayCtl.RsClient)
	if err != nil {
		t.Errorf("Failed%s", err)
	}
	xrayCtl.CmdConn.Close()

}
