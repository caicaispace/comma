package main

import (
	"log"

	gatewayGrpc "comma/pkg/server/grpc/gateway"

	"github.com/caicaispace/gohelper/server"
	"golang.org/x/sync/errgroup"
)

var g errgroup.Group

// func init() {
// 	// db.New(&setting.DBSetting{
// 	// 	Host:     "127.0.0.1",
// 	// 	Port:     "3306",
// 	// 	Username: "root",
// 	// 	Password: "123456",
// 	// 	DbName:   "comma",
// 	// })
// }

// func httpServerStart(serverAddr string) error {
// 	s := httpServer.NewServer()
// 	s.SetServerAddr(serverAddr)
// 	// s.UseTrace(TRACE_URL, "gateway", serverAddr)
// 	gatewayHttp.NewServer(s)
// 	s.Start()
// 	return nil
// }

// func jsonRpcServerStart(serverAddr string) error {
// 	s := jsonrpcServer.NewServer()
// 	s.SetServerAddr(serverAddr)
// 	s.RegisterService(gatewayJsonRpc.GetInstance())
// 	s.Start()
// 	return nil
// }

func grpcServerStart(serverAddr string) error {
	s := gatewayGrpc.NewServer()
	s.SetServerAddr(serverAddr)
	s.Start()
	return nil
}

func main() {
	server.New()
	// g.Go(func() error {
	// 	return jsonRpcServerStart("127.0.0.1:3231")
	// })
	// g.Go(func() error {
	// 	return httpServerStart("127.0.0.1:3232")
	// })
	g.Go(func() error {
		return grpcServerStart("127.0.0.1:3232")
	})
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
