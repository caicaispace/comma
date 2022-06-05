package main

import (
	"embed"
	"goaway/pkg/library/db"
	serverSetting "goaway/pkg/library/net"
	httpServer "goaway/pkg/library/net/http"
	"goaway/pkg/library/setting"
	"goaway/pkg/library/util/config"
	"goaway/pkg/model"
	adminServer "goaway/pkg/server/http/admin"
	gatewayServer "goaway/pkg/server/http/gateway"
	bannedService "goaway/pkg/service/banned"
	"io/fs"
	"net/http"

	//bannedJsonRpc `goaway/pkg/service/banned/server/jsonrpc`

	jsonrpcServer "goaway/pkg/library/net/jsonrpc"
	gatewayJsonRpc "goaway/pkg/server/jsonrpc/gateway"
	segmentJsonRpc "goaway/pkg/server/jsonrpc/segment"
	"log"

	"golang.org/x/sync/errgroup"
)

const (
	TRACE_URL = ""
)

type Service struct {
	IsOpen bool
}

var (
	//go:embed static/*
	static   embed.FS
	g        errgroup.Group
	services = map[string]*Service{
		"admin":   {IsOpen: true},
		"jsonrpc": {IsOpen: true},
		"grpc":    {IsOpen: true},
	}
)

func beforeStart() {
	if config.GetInstance().GetEnv() == "dev" {
		db.NewWithAddr(config.GetInstance().GetDB())
	}
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
	adminServer.NewServer(s.Engine)
	fe, _ := fs.Sub(static, "static")
	s.Engine.StaticFS("ui", http.FS(fe))
	s.Start()
	return nil
}

func gatewayServerStart(serverAddr string) error {
	// s := httpServer.NewServer()
	// s.SetServerAddr(serverAddr)
	// s.SetTrace(TRACE_URL, "gateway", serverAddr)
	// gatewayServer.NewServer(s)
	// s.Start()
	gatewayServer.NewServer(serverAddr)
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
	g.Go(func() error {
		// gateway server
		return gatewayServerStart(config.GetInstance().GetServerHost() + ":9400")
		return nil
	})
	g.Go(func() error {
		// jsonrpc service
		if services["jsonrpc"].IsOpen {
			return jsonRpcServerStart(config.GetInstance().GetServerHost() + ":9401")
		}
		return nil
	})
	g.Go(func() error {
		// admin server
		if services["admin"].IsOpen {
			return adminServerStart(config.GetInstance().GetServerHost() + ":9402")
		}
		return nil
	})
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
