package main

import (
	"log"

	"goaway/pkg/demo/jsonrpc/jsonrpc"
	serverSetting "goaway/pkg/library/net"
	"goaway/pkg/library/net/http"
	jsonRpcServer "goaway/pkg/library/net/jsonrpc"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

var g errgroup.Group

// https://www.jianshu.com/p/74ac2439afb2

func httpServerStart(serverAddr string) error {
	s := http.NewServer()
	s.SetServerAddr(serverAddr)
	apiV1 := s.Engine.Group("/v1")
	{
		apiV1.GET("/test", func(c *gin.Context) {
			ctx := http.Context{C: c}
			ctx.Success(gin.H{
				"good": "lock",
			}, nil)
		})
	}
	s.Start()
	return nil
}

func jsonRpcServerStart(serverAddr string) error {
	s := jsonRpcServer.NewServer()
	s.SetServerAddr(serverAddr)
	s.RegisterService(new(jsonrpc.Demo))
	s.Start()
	return nil
}

func main() {
	serverSetting.New()
	g.Go(func() error {
		return jsonRpcServerStart("127.0.0.1:8081")
	})
	g.Go(func() error {
		return httpServerStart("127.0.0.1:8082")
	})
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
	//go jsonRpcServerStart("127.0.0.1:8081")
	//go httpServerStart("127.0.0.1:8082")
	////阻塞程序
	//select {}
}
