package examples

import (
	"context"
	statsService "github.com/xtls/xray-core/app/stats/command"
)

/*
	获取运行数据, 如下
	NumGoroutine:17  NumGC:2  Alloc:1711192  TotalAlloc:2359880  Sys:14440840  Mallocs:19101  Frees:7242  LiveObjects:11859  PauseTotalNs:4983200  Uptime:31
*/
func getSysStats(c statsService.StatsServiceClient) (stats *statsService.SysStatsResponse, err error) {
	stats, err = c.GetSysStats(context.Background(), &statsService.SysStatsRequest{})

	return
}
