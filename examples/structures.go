package examples

import (
	loggerService "github.com/xtls/xray-core/app/log/command"
	handlerService "github.com/xtls/xray-core/app/proxyman/command"
	routingService "github.com/xtls/xray-core/app/router/command"
	statsService "github.com/xtls/xray-core/app/stats/command"
	"google.golang.org/grpc"
)

type BaseConfig struct {
	APIAddress string
	APIPort    uint16
}

type XrayController struct {
	HsClient handlerService.HandlerServiceClient
	SsClient statsService.StatsServiceClient
	LsClient loggerService.LoggerServiceClient
	RsClient routingService.RoutingServiceClient
	CmdConn  *grpc.ClientConn
}
