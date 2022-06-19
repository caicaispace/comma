package main

import (
	"flag"
	"fmt"
	"time"

	"goaway/pkg/library/db"
	serverSetting "goaway/pkg/library/net"
	jsonrpcServer "goaway/pkg/library/net/jsonrpc"
	"goaway/pkg/library/setting"
	"goaway/pkg/library/util/metric"
	"goaway/pkg/library/util/task"
	bannedJsonRpc "goaway/pkg/server/jsonrpc/banned"
)

var (
	// metric
	metricJob          = flag.String("metric-job", "goaway", "prometheus job name")
	metricInstance     = flag.String("metric-instance", "", "prometheus instance name")
	metricAddress      = flag.String("metric-address", "127.0.0.1:9091", "prometheus proxy address")
	metricIntervalSync = flag.Uint64("interval-metric-sync", 1, "Interval(sec): metric sync")
)

func init() {
	config := &setting.DBSetting{
		Host:     "127.0.0.1",
		Port:     "3306",
		Username: "root",
		Password: "123456",
		DbName:   "goaway",
	}
	db.New(config)
}

func main() {
	serverSetting.New()
	s := jsonrpcServer.NewServer()
	s.SetServerAddr("127.0.0.1:3231")
	s.RegisterService(bannedJsonRpc.GetInstance())
	s.SetServerAfterStartFunc(func() {
		conf := metric.NewMetricCfg(*metricJob, *metricInstance, *metricAddress, time.Second*time.Duration(*metricIntervalSync))
		runner := task.NewRunner()
		metric.StartMetricsPush(runner, conf)
		m := metric.NewMetric()
		ticker := time.NewTicker(1 * time.Second)
		go func(t *time.Ticker) {
			for {
				<-t.C
				m.PostRequest("test", true, time.Time{})
				fmt.Println("get ticker", time.Now().Format("2006-01-02 15:04:05"))
			}
		}(ticker)
	})
	s.Start()
}
