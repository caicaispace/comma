package main

import (
	"embed"
	"flag"
	"goaway/pkg/library/db"
	serverSetting "goaway/pkg/library/net"
	httpServer "goaway/pkg/library/net/http"
	"goaway/pkg/library/setting"
	"goaway/pkg/library/util/config"
	"goaway/pkg/model"
	adminHttpServer "goaway/pkg/server/http/admin"
	gatewayHttpServer "goaway/pkg/server/http/gateway"
	"io/fs"
	"net/http"

	jsonrpcServer "goaway/pkg/library/net/jsonrpc"
	gatewayJsonRpc "goaway/pkg/server/jsonrpc/gateway"
	segmentJsonRpc "goaway/pkg/server/jsonrpc/segment"

	//bannedJsonRpc `goaway/pkg/service/banned/server/jsonrpc`
	bannedService "goaway/pkg/service/banned"
	"log"

	"golang.org/x/sync/errgroup"
)

var (
	// metric
	metricJob          = flag.String("metric-job", "", "prometheus job name")
	metricInstance     = flag.String("metric-instance", "", "prometheus instance name")
	metricAddress      = flag.String("metric-address", "", "prometheus proxy address")
	metricIntervalSync = flag.Uint64("interval-metric-sync", 0, "Interval(sec): metric sync")
)

const (
	TRACE_URL = ""
)

type Service struct {
	AdminIsOpen   bool
	JsonrpcIsOpen bool
	GrpcIsOpen    bool
}

var (
	//go:embed static/*
	static   embed.FS
	group    errgroup.Group
	services = &Service{
		AdminIsOpen:   true,
		JsonrpcIsOpen: true,
		GrpcIsOpen:    false,
	}
)

func beforeStart() {
	if config.GetInstance().GetEnv() == "dev" {
		db.NewWithAddr(config.GetInstance().GetDB())
	}
	// metric.NewMetricCfg(*metricJob, *metricInstance, *metricAddress, time.Second*time.Duration(*metricIntervalSync))
	if setting.Database.AutoMigrate == true {
		db.DB().AutoMigrate(
			&model.DictBanned{},
			&model.DictFestival{},
			&model.DictHighFrequency{},
			&model.DictHyponym{},
			&model.DictPinyin{},
			&model.DictProject{},
			&model.DictRedirect{},
			&model.DictStop{},
			&model.DictSynonyms{},
			&model.DictVersion{},
			&model.DictWeight{},
			&model.DictWord{},
		)
	}
	loadService()
}

func loadService() {
	bannedService.GetInstance()
}

func adminServerStart(serverAddr string) error {
	s := httpServer.NewServer()
	s.SetServerAddr(serverAddr)
	// s.UseTrace(TRACE_URL, "goaway-admin", serverAddr)
	// s.UseGrafana()
	adminHttpServer.NewServer(s.Engine)
	fe, _ := fs.Sub(static, "static")
	s.Engine.StaticFS("ui", http.FS(fe))
	s.Start()
	return nil
}

func gatewayServerStart(serverAddr string) error {
	s := httpServer.NewServer()
	s.SetServerAddr(serverAddr)
	// s.UseTrace(TRACE_URL, "gateway", serverAddr)
	gatewayHttpServer.NewServer(s)
	s.Start()
	return nil
}

func jsonRpcServerStart(serverAddr string) error {
	s := jsonrpcServer.NewServer()
	s.SetServerAddr(serverAddr)
	// s.RegisterService(bannedJsonRpc.GetInstance())
	s.RegisterService(segmentJsonRpc.GetInstance())
	s.RegisterService(gatewayJsonRpc.GetInstance())
	s.Start()
	return nil
}

func main() {
	serverSetting.New()
	beforeStart()
	group.Go(func() error {
		// gateway server
		return gatewayServerStart(config.GetInstance().GetServerHost() + ":9400")
	})
	group.Go(func() error {
		// jsonrpc service
		if services.JsonrpcIsOpen {
			return jsonRpcServerStart(config.GetInstance().GetServerHost() + ":9401")
		}
		return nil
	})
	group.Go(func() error {
		// admin server
		if services.AdminIsOpen {
			return adminServerStart(config.GetInstance().GetServerHost() + ":9402")
		}
		return nil
	})
	if err := group.Wait(); err != nil {
		log.Fatal(err)
	}
}
