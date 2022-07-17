package main

import (
	"log"

	"comma/pkg/library/db"
	serverSetting "comma/pkg/library/net"
	httpServer "comma/pkg/library/net/http"
	jsonrpcServer "comma/pkg/library/net/jsonrpc"
	"comma/pkg/library/setting"
	gatewayServer "comma/pkg/server/http/gateway"
	gatewayJsonRpc "comma/pkg/server/jsonrpc/gateway"

	"golang.org/x/sync/errgroup"
)

var g errgroup.Group

func init() {
	db.New(&setting.DBSetting{
		Host:     "127.0.0.1",
		Port:     "3306",
		Username: "root",
		Password: "123456",
		DbName:   "comma",
	})
}

func httpServerStart(serverAddr string) error {
	s := httpServer.NewServer()
	s.SetServerAddr(serverAddr)
	// s.UseTrace(TRACE_URL, "gateway", serverAddr)
	gatewayServer.NewServer(s)
	s.Start()
	return nil
}

func jsonRpcServerStart(serverAddr string) error {
	s := jsonrpcServer.NewServer()
	s.SetServerAddr(serverAddr)
	s.RegisterService(gatewayJsonRpc.GetInstance())
	s.Start()
	return nil
}

func main() {
	serverSetting.New()
	g.Go(func() error {
		return jsonRpcServerStart("127.0.0.1:3231")
	})
	g.Go(func() error {
		return httpServerStart("127.0.0.1:3232")
	})
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
