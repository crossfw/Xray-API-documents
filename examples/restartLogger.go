package examples

import (
	"context"
	loggerService "github.com/xtls/xray-core/app/log/command"
)

func restartLogger(c loggerService.LoggerServiceClient) (err error) {
	_, err = c.RestartLogger(context.Background(), &loggerService.RestartLoggerRequest{})

	return
}
