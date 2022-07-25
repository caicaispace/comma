package main

import (
	"comma/pkg/library/db"
	"comma/pkg/library/setting"
	segmentJsonRpc "comma/pkg/server/jsonrpc/segment"

	jsonrpcServer "github.com/caicaispace/gohelper/server/jsonrpc"
)

func init() {
	db.New(&setting.DBSetting{
		Host:     "127.0.0.1",
		Port:     "3306",
		Username: "root",
		Password: "123456",
		DbName:   "comma",
	})
}

func main() {
	s := jsonrpcServer.NewServer()
	s.SetServerAddr("127.0.0.1:3231")
	s.RegisterService(segmentJsonRpc.GetInstance())
	s.Start()
}
