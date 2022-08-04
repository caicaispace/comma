package main

import (
	segmentJsonRpc "comma/pkg/server/jsonrpc/segment"

	"github.com/caicaispace/gohelper/orm/gorm"
	jsonrpcServer "github.com/caicaispace/gohelper/server/jsonrpc"
	"github.com/caicaispace/gohelper/setting"
)

func init() {
	config := &setting.DBSetting{
		Host:     "127.0.0.1",
		Port:     "3306",
		Username: "root",
		Password: "123456",
		DbName:   "comma",
	}
	gorm.New().AddConnWithConfig(config, connName).GetDB(connName)
}

func main() {
	s := jsonrpcServer.NewServer()
	s.SetServerAddr("127.0.0.1:3231")
	s.RegisterService(segmentJsonRpc.GetInstance())
	s.Start()
}
