package main

import (
	"comma/pkg/library/db"
	"comma/pkg/library/setting"
	bannedJsonRpc "comma/pkg/server/jsonrpc/banned"

	jsonrpcServer "github.com/caicaispace/gohelper/server/jsonrpc"
)

func init() {
	config := &setting.DBSetting{
		Host:     "127.0.0.1",
		Port:     "3306",
		Username: "root",
		Password: "123456",
		DbName:   "comma",
	}
	db.New(config)
}

func main() {
	s := jsonrpcServer.NewServer()
	s.SetServerAddr("127.0.0.1:3231")
	s.RegisterService(bannedJsonRpc.GetInstance())
	s.Start()
}
