package examples

import (
	"context"
	"github.com/xtls/xray-core/app/router/command"
)

func addRouting(client command.RoutingServiceClient) error {
	_, err := client.TestRoute(context.Background(), &command.TestRouteRequest{
		RoutingContext: &command.RoutingContext{
			InboundTag:  "p0",
			OutboundTag: "block",
		},
		FieldSelectors: nil,
		PublishResult:  false,
	})

	return err
}
