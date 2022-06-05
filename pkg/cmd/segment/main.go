package main

import (
	"goaway/pkg/library/db"
	jsonrpcServer "goaway/pkg/library/net/jsonrpc"
	"goaway/pkg/library/setting"
	segmentJsonRpc "goaway/pkg/server/jsonrpc/segment"
)

func init() {
	db.New(&setting.DBSetting{
		Host:     "127.0.0.1",
		Port:     "3306",
		Username: "root",
		Password: "123456",
		DbName:   "goaway",
	})
}

func main() {
	s := jsonrpcServer.NewServer()
	s.SetServerAddr("127.0.0.1:3231")
	s.RegisterService(segmentJsonRpc.GetInstance())
	s.Start()
}
