package examples

import (
	"errors"
	"fmt"
	loggerService "github.com/xtls/xray-core/app/log/command"
	handlerService "github.com/xtls/xray-core/app/proxyman/command"
	routingService "github.com/xtls/xray-core/app/router/command"
	statsService "github.com/xtls/xray-core/app/stats/command"
	"google.golang.org/grpc"
)

func (xrayCtl *XrayController) Init(cfg *BaseConfig) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("init Xray API error - %s", r))
		}
	}()
	xrayCtl.CmdConn, err = grpc.Dial(fmt.Sprintf("%s:%d", cfg.APIAddress, cfg.APIPort), grpc.WithInsecure())
	if err != nil {
		return err
	}

	xrayCtl.HsClient = handlerService.NewHandlerServiceClient(xrayCtl.CmdConn)
	xrayCtl.SsClient = statsService.NewStatsServiceClient(xrayCtl.CmdConn)
	xrayCtl.LsClient = loggerService.NewLoggerServiceClient(xrayCtl.CmdConn)
	//Not implement
	xrayCtl.RsClient = routingService.NewRoutingServiceClient(xrayCtl.CmdConn)

	return
}
