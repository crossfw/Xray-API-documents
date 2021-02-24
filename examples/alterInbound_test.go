package examples

import "testing"

func TestAlertInbound(t *testing.T) {
	var (
		xrayCtl *XrayController
		cfg     = &BaseConfig{
			APIAddress: "127.0.0.1",
			APIPort:    10085,
		}
		user = UserInfo{
			Uuid:       "10354ac4-9ec1-4864-ba3e-f5fd35869ef8",
			AlertId:    0,
			Level:      0,
			InTag:      "proxy0",
			Email:      "love@xray.com",
			CipherType: "aes-256-gcm",
			Password:   "xrayisthebest",
		}
	)
	xrayCtl = new(XrayController)
	err := xrayCtl.Init(cfg)
	defer xrayCtl.CmdConn.Close()
	if err != nil {
		t.Errorf("Failed %s", err)
	}
	err = addVmessUser(xrayCtl.HsClient, &user)
	if err != nil {
		t.Errorf("Failed %s", err)
	}

}
