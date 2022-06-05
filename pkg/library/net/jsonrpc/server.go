package jsonrpc

import (
	"errors"
	"net/http"

	"goaway/pkg/library/net/jsonrpc/common"
	"goaway/pkg/library/net/jsonrpc/server"
)

type Error common.Error

type ServerInterface interface {
	SetHttpRouters(routersMap *map[string]func(w http.ResponseWriter, r *http.Request))
	SetBeforeFunc(func(id interface{}, method string, params interface{}) error)
	SetAfterFunc(func(id interface{}, method string, result interface{}) error)
	SetServerBeforeStartFunc(func())
	SetServerAfterStartFunc(func())
	SetOptions(interface{})
	SetRateLimit(float64, int64)
	Start()
	Register(s interface{})
}

func New(protocol string, ip string, port string) (ServerInterface, error) {
	var err error
	switch protocol {
	case "http":
		return server.NewHttpServer(ip, port), err
	case "tcp":
		return server.NewTcpServer(ip, port), err
	}
	return nil, errors.New("The protocol can not be supported")
}
