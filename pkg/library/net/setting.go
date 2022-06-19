package net

import (
	"flag"
	"time"

	"goaway/pkg/library/core/l"
	"goaway/pkg/library/setting"
	util2 "goaway/pkg/library/util"
)

var (
	// server
	host         = flag.String("host", "127.0.0.1", "server address host")
	port         = flag.String("port", "8081", "server address port")
	addr         = flag.String("addr", "127.0.0.1:8081", "Server: client http http entry point")
	env          = flag.String("env", "prod", "server environment variable")
	readTimeout  = flag.Int64("rt", 60, "Server: client http read timeout")
	writeTimeout = flag.Int64("wt", 60, "Server: client http write timeout")
	// common
	timeFormat = flag.String("time-format", "20060102", "App: time format")
	// log
	logPath      = flag.String("log-path", "logs/", "App: log file path")
	logPrefix    = flag.String("log-pref", "log_", "App: log file prefix")
	logExtension = flag.String("log-ext", "log", "App: log file extension")
	// database
	autoMigrate = flag.Bool("at", false, "auto migrate run auto migration for given models")
	// metric
	metricEnable       = flag.Bool("metric-enable", true, "prometheus is enable")
	metricJob          = flag.String("metric-job", "goaway", "prometheus job name")
	metricInstance     = flag.String("metric-instance", "", "prometheus instance name")
	metricAddress      = flag.String("metric-address", "127.0.0.1:9091", "prometheus proxy address")
	metricIntervalSync = flag.Uint64("interval-metric-sync", 1, "Interval(sec): metric sync")
)

func init() {
	flag.Parse()
	initServerSetting()
	initAppSetting()
	initDBSetting()
	initMetricSetting()
}

func New() {
	l.Setup()
}

func initServerSetting() {
	setting.Server.Env = *env
	setting.Server.Host = util2.If(*env == "dev", *host, util2.LocalIP()).(string)
	setting.Server.Port = *port
	setting.Server.Addr = setting.Server.Host + ":" + *port
	setting.Server.RootPath = util2.GetCurrentAbPath()
	setting.Server.ReadTimeout = time.Duration(*readTimeout)
	setting.Server.WriteTimeout = time.Duration(*writeTimeout)
	// if setting.ServerSetting.Env == "dev" && strings.Contains(setting.ServerSetting.RootPath, "banned") == false {
	// 	fmt.Println("------------------ Please start the service in the app directory ------------------")
	// 	os.Exit(0)
	// }
}

func initAppSetting() {
	setting.App.TimeFormat = *timeFormat
	setting.App.RootPath = util2.GetCurrentAbPath()
	setting.App.LogPath = *logPath
	setting.App.LogPrefix = *logPrefix
	setting.App.LogExtension = *logExtension
}

func initDBSetting() {
	setting.Database.AutoMigrate = *autoMigrate
}

func initMetricSetting() {
	setting.Metric.Enable = *metricEnable
	setting.Metric.Job = *metricJob
	setting.Metric.Instance = *metricInstance
	setting.Metric.Address = *metricAddress
	setting.Metric.IntervalSync = *metricIntervalSync
}
