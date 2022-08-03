package main

import (
	"comma/pkg/demo/jsonrpc/jsonrpc"

	_ "github.com/caicaispace/gohelper/server"
	jsonrpcServer "github.com/caicaispace/gohelper/server/jsonrpc"
)

// https://www.jianshu.com/p/74ac2439afb2

func serverStart(addr string) {
	s := jsonrpcServer.NewServer()
	s.SetServerAddr(addr)
	s.RegisterService(new(jsonrpc.Demo))
	s.Start()
}

func main() {
	serverStart("127.0.0.1:8081")
}
