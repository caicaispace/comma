package main

import (
	"goaway/pkg/library/db"
	jsonrpcServer "goaway/pkg/library/net/jsonrpc"
	"goaway/pkg/library/setting"
	bannedJsonRpc "goaway/pkg/server/jsonrpc/banned"
)

func init() {
	config := &setting.DBSetting{
		Host:     "127.0.0.1",
		Port:     "3306",
		Username: "root",
		Password: "123456",
		DbName:   "goaway",
	}
	db.New(config)
}

func main() {
	s := jsonrpcServer.NewServer()
	s.SetServerAddr("127.0.0.1:3231")
	s.RegisterService(bannedJsonRpc.GetInstance())
	s.Start()
}
