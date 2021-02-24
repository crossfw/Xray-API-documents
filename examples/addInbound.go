package examples

import (
	"context"
	"github.com/xtls/xray-core/app/proxyman"
	"github.com/xtls/xray-core/app/proxyman/command"
	"github.com/xtls/xray-core/common/net"
	"github.com/xtls/xray-core/common/protocol/tls/cert"
	"github.com/xtls/xray-core/common/serial"
	"github.com/xtls/xray-core/core"
	_ "github.com/xtls/xray-core/infra/conf"
	//ssInbound "github.com/xtls/xray-core/proxy/shadowsocks"
	//trojanInbound "github.com/xtls/xray-core/proxy/trojan"
	vmessInbound "github.com/xtls/xray-core/proxy/vmess/inbound"
	"github.com/xtls/xray-core/transport/internet"
	_ "github.com/xtls/xray-core/transport/internet/tcp"
	"github.com/xtls/xray-core/transport/internet/tls"
	"github.com/xtls/xray-core/transport/internet/websocket"
)

func addInbound(client command.HandlerServiceClient) error {

	_, err := client.AddInbound(context.Background(), &command.AddInboundRequest{
		Inbound: &core.InboundHandlerConfig{
			Tag: "proxy0",
			ReceiverSettings: serial.ToTypedMessage(&proxyman.ReceiverConfig{
				// 监听端口 12345
				PortRange: net.SinglePortRange(net.Port(12345)),
				// 监听地址, 默认0.0.0.0
				Listen: net.NewIPOrDomain(net.AnyIP),
				// 流量探测
				SniffingSettings: &proxyman.SniffingConfig{
					Enabled:             true,
					DestinationOverride: []string{"http", "tls"},
				},
				// 传输方式
				StreamSettings: &internet.StreamConfig{
					/*
						传输方式名称
						请自行在 github.com/xtls/xray-core/transport/internet/config.pb.go 中寻找支持的协议
						截至1.3.0 目前支持
						"TCP",
						"UDP",
						"MKCP",
						"WebSocket",
						"HTTP",
						"DomainSocket",
					*/
					ProtocolName: "WebSocket",
					TransportSettings: []*internet.TransportConfig{
						{
							ProtocolName: "WebSocket",
							/*
								选定传输方式后,请去 github.com/xtls/xray-core/transport/internet 下你选定方式的文件夹中导入config结构
								如选定WebSocket则需要使用 github.com/xtls/xray-core/transport/internet/websocket/config.pb.go 中的 Config struct
								结构内容请自行翻看代码(Ctrl + 左键)
							*/
							Settings: serial.ToTypedMessage(&websocket.Config{
								Path: "/web",
								Header: []*websocket.Header{
									{
										Key:   "Host",
										Value: "www.xray.best",
									},
								},
								AcceptProxyProtocol: false,
							},
							),
						},
					},
					/*
						传输层加密
						请在 github.com/xtls/xray-core/transport/internet/ 中选择合适的传输层加密方式
						截至1.3.0 目前支持
						TLS
						XTLS
						留空即为None
					*/
					SecurityType: serial.GetMessageType(&tls.Config{}),
					SecuritySettings: []*serial.TypedMessage{
						serial.ToTypedMessage(&tls.Config{
							//Auto build
							Certificate: []*tls.Certificate{tls.ParseCertificate(cert.MustGenerate(nil))},
						}),
					},
				},
			}),
			ProxySettings: serial.ToTypedMessage(&vmessInbound.Config{}),
		},
	})

	return err
}
