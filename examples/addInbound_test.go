package examples

import "testing"

func TestAddInbound(t *testing.T) {
	// 先指定 API 端口和地址
	var (
		xrayCtl *XrayController
		cfg     = &BaseConfig{
			APIAddress: "127.0.0.1",
			APIPort:    10085,
		}
	)
	// 初始化 Clients
	xrayCtl = new(XrayController)
	err := xrayCtl.Init(cfg)
	defer xrayCtl.CmdConn.Close()
	if err != nil {
		t.Errorf("Failed %s", err)
	}
	// 此处为执行命令部分
	err = addInbound(xrayCtl.HsClient)
	if err != nil {
		t.Errorf("Failed %s", err)
	}

}
