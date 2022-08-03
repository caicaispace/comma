package gateway

import (
	gatewayServer "comma/pkg/server/grpc/gateway"
	gatewayService "comma/pkg/service/gateway"
	"fmt"
	"time"

	"github.com/caicaispace/gohelper/cluster/etcd"
	"github.com/caicaispace/gohelper/server/grpc/server"
	clientV3 "go.etcd.io/etcd/client/v3"
)

type Gateway struct {
	GatewayService *gatewayService.Service
	serverAddr     string
}

func NewServer() *Gateway {
	return &Gateway{
		// GatewayService: gatewayService.GetInstance(),
	}
}

func (that *Gateway) SetServerAddr(addr string) {
	that.serverAddr = addr
}

func (that *Gateway) Start() {
	register, _ := etcd.NewRegister(&etcd.NodeInfo{
		Addr:     that.serverAddr,
		Name:     "comma",
		UniqueId: fmt.Sprintf("discovery/comma/instance_id/%s", "888"),
	}, clientV3.Config{
		Endpoints:            []string{"127.0.0.1:2379"},
		DialTimeout:          2 * time.Second,
		DialKeepAliveTime:    time.Second,
		DialKeepAliveTimeout: time.Second,
	})
	go register.Run()
	s := server.NewServer(that.serverAddr)
	gatewayServer.RegisterGatewayServer(s.GrpcServer, &gatewayServer.Service{})
	s.Start()
}
