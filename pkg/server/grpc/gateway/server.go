package gateway

import (
	gatewayService "comma/pkg/service/gateway"
	context "context"
	"encoding/json"
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
		Name:     "gateway",
		UniqueId: fmt.Sprintf("discovery/all/instance_id/%s", "gateway"),
	}, clientV3.Config{
		Endpoints:            []string{"127.0.0.1:2379"},
		DialTimeout:          2 * time.Second,
		DialKeepAliveTime:    time.Second,
		DialKeepAliveTimeout: time.Second,
	})
	go register.Run()
	s := server.New(that.serverAddr)
	RegisterGatewayServer(s.GrpcServer, that)
	s.Start()
}

func (s *Gateway) List(c context.Context, in *ListReq) (*ListRsp, error) {
	return &ListRsp{List: "user list"}, nil
}

func (*Gateway) Search(c context.Context, in *SearchReq) (*SearchRsp, error) {
	esData, err := gatewayService.GetInstance().DispatchWithJsonRpc(in.Index, in.Type, in.Body, "search")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	var out string
	json.Unmarshal(esData, &out)
	return &SearchRsp{Data: out}, nil
}
