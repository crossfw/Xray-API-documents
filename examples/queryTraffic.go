package examples

import (
	"context"
	statsService "github.com/xtls/xray-core/app/stats/command"
)

/*
	先按照 https://xtls.github.io/config/base/policy/ 打开入/出 站或用户流量统计
	请参照 https://xtls.github.io/config/base/stats/ 来生成一条查询语句， 例如
	“user>>>love@xray.com>>>traffic>>>uplink”	查询 email 为 love@xray.com 的用户在所有的入站中的上行流量
	可以选择在查询完之后重置流量信息。

	目前支持 User, Inbound, Outbound 的上下行流量查询
*/

func queryTraffic(c statsService.StatsServiceClient, ptn string, reset bool) (traffic int64, err error) {
	// 如果查无此用户或 bound 则返回-1, 默认值 -1
	traffic = -1
	resp, err := c.QueryStats(context.Background(), &statsService.QueryStatsRequest{
		// 这里是查询语句，例如 “user>>>love@xray.com>>>traffic>>>uplink” 表示查询用户 email 为 love@xray.com 在所有入站中的上行流量
		Pattern: ptn,
		// 是否重置流量信息(true, false)，即完成查询后是否把流量统计归零
		Reset_: reset, // reset traffic data everytime
	})
	if err != nil {
		return
	}
	// Get traffic data
	stat := resp.GetStat()
	// 判断返回 是否成功
	// 返回样例，value 值是我们需要的: [name:"inbound>>>proxy0>>>traffic>>>downlink" value:348789]
	if len(stat) != 0 {
		// 返回流量数据 byte
		traffic = stat[0].Value
	}

	return
}
