package examples

import (
	"context"
	"github.com/xtls/xray-core/app/proxyman/command"
)

// 使用 Tag 操作，非常简单
func removeInbound(client command.HandlerServiceClient, tag string) error {
	_, err := client.RemoveInbound(context.Background(), &command.RemoveInboundRequest{
		Tag: tag,
	})
	return err
}
