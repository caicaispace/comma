package main

import (
	httpServer "github.com/caicaispace/gohelper/server/http"

	"comma/pkg/model"
	"embed"
	"io/fs"
	"log"
	"net/http"
	"time"

	adminHttpServer "comma/pkg/server/http/admin"
	gatewayHttpServer "comma/pkg/server/http/gateway"
	gatewayJsonRpc "comma/pkg/server/jsonrpc/gateway"
	segmentJsonRpc "comma/pkg/server/jsonrpc/segment"

	jsonrpcServer "github.com/caicaispace/gohelper/server/jsonrpc"

	//bannedJsonRpc `comma/pkg/service/banned/server/jsonrpc`
	bannedService "comma/pkg/service/banned"

	"github.com/caicaispace/gohelper/config"
	"github.com/caicaispace/gohelper/metric"
	"github.com/caicaispace/gohelper/orm/gorm"
	"github.com/caicaispace/gohelper/server"
	"github.com/caicaispace/gohelper/setting"
	"github.com/caicaispace/gohelper/task"
	"golang.org/x/sync/errgroup"
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
		gorm.GetInstance().AddConnWithDns(config.GetInstance().GetDbDns(), "")
	}
	if setting.Database.AutoMigrate {
		gorm.GetInstance().GetDB("").AutoMigrate(
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
	if config.GetInstance().GetMetricIsEnable() {
		conf := config.GetInstance().GetMetric()
		metric.StartMetricsPush(task.NewRunner(), metric.NewMetricCfg(
			conf.Job,
			conf.Instance,
			conf.Address,
			time.Second*time.Duration(conf.IntervalSync),
		))
	}
	loadService()
}

func loadService() {
	bannedService.GetInstance()
}

func adminServerStart(serverAddr string) error {
	s := httpServer.NewServer()
	s.AddServerAddr(serverAddr)
	// s.UseTrace(TRACE_URL, "comma-admin", serverAddr)
	// s.UseGrafana()
	adminHttpServer.NewServer(s.Engine)
	fe, _ := fs.Sub(static, "static")
	s.Engine.StaticFS("ui", http.FS(fe))
	s.Start()
	return nil
}

func gatewayServerStart(serverAddr string) error {
	s := httpServer.NewServer()
	s.AddServerAddr(serverAddr)
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
	server.New()
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
