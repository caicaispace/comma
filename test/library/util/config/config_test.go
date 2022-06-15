package config

import (
	"fmt"
	"goaway/pkg/library/util/config"
	"testing"
)

var conf = config.GetInstance()

func TestConf_New(t *testing.T) {
	fmt.Printf("title: %+v\n\n", conf.Title)
	fmt.Printf("Server Config: %+v\n\n", conf.Server)
	fmt.Printf("Database Config: %+v\n\n", conf.DB)
	fmt.Printf("ElasticSearch Config: %+v\n\n", conf.ES)
}

func TestConf_GetServerHost(t *testing.T) {
	t.Log(conf.GetServerHost())
}

func TestConf_GetEs(t *testing.T) {
	t.Log(conf.GetEs())
}

func TestConf_GetEsRoute(t *testing.T) {
	t.Log(conf.GetEsRoute("search_all_v2", "all"))
}

func TestConf_GetEsFilterType(t *testing.T) {
	t.Log(conf.GetEsFilterType("search_all_v2", "all"))
}

func TestConf_GetEsProjectId(t *testing.T) {
	t.Log(conf.GetEsProjectId("search_all_v2", "all"))
}

func TestConf_GetProxyRoutes(t *testing.T) {
	t.Log(conf.GetProxyRoutes())
}
