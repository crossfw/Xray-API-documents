package examples

import (
	"fmt"
	loggerService "github.com/xtls/xray-core/app/log/command"
	handlerService "github.com/xtls/xray-core/app/proxyman/command"
	routingService "github.com/xtls/xray-core/app/router/command"
	statsService "github.com/xtls/xray-core/app/stats/command"
	"google.golang.org/grpc"
)

// 取得API操作的Client

func (xrayCtl *XrayController) Init(cfg *BaseConfig) (err error) {
	// 先取得ClientConn, 用完记得close
	xrayCtl.CmdConn, err = grpc.Dial(fmt.Sprintf("%s:%d", cfg.APIAddress, cfg.APIPort), grpc.WithInsecure())
	if err != nil {
		return err
	}

	// 依次获取API Client, 可根据需求删减
	xrayCtl.HsClient = handlerService.NewHandlerServiceClient(xrayCtl.CmdConn)
	xrayCtl.SsClient = statsService.NewStatsServiceClient(xrayCtl.CmdConn)
	xrayCtl.LsClient = loggerService.NewLoggerServiceClient(xrayCtl.CmdConn)
	//Not implement
	xrayCtl.RsClient = routingService.NewRoutingServiceClient(xrayCtl.CmdConn)

	return
}
