package config

import (
	"encoding/xml"
	"io/ioutil"
	"sync"

	"comma/pkg/library/setting"
	"comma/pkg/library/util"
)

type database struct {
	Dict string `xml:"dict"`
}

type elasticsearch struct {
	IP       string `xml:"ip"`
	User     string `xml:"user"`
	Password string `xml:"password"`
}

type route struct {
	IndexName            string `xml:"indexName"`
	TypeName             string `xml:"typeName"`
	DestinationIndexName string `xml:"destinationIndexName"`
	DestinationTypeName  string `xml:"destinationTypeName"`
	DestinationIP        string `xml:"destinationIp"`
	DestinationPort      string `xml:"destinationPort"`
	FilterType           int    `xml:"filterType"`
	ProjectID            int    `xml:"projectId"`
}

type routes struct {
	Route []route `xml:"Route"`
}

type mappings struct {
	RouteMap      map[string]string
	ProxyRoutes   []string
	FilterTypeMap map[string]int
	ProjectMap    map[string]int
}

type configXml struct {
	Database      database      `xml:"Database"`
	Elasticsearch elasticsearch `xml:"Elasticsearch"`
	Routes        routes        `xml:"Routes"`
	Mappings      mappings
}

var (
	cx     *configXml
	cxOnce sync.Once
)

func GetXmlInstance() *configXml {
	cxOnce.Do(func() {
		mp := mappings{
			RouteMap:      make(map[string]string),
			FilterTypeMap: make(map[string]int),
			ProjectMap:    make(map[string]int),
		}
		cx = &configXml{
			Mappings: mp,
		}
		cx.loadXmlFile()
	})
	return cx
}

func (cx *configXml) GetEs() *elasticsearch {
	return &cx.Elasticsearch
}

func (*configXml) GetEsRoute(indexName string, typeName string) string {
	key, exist := cx.Mappings.RouteMap[indexName+"-"+typeName]
	return util.If(exist == false, "", key).(string)
}

func (*configXml) GetProxyRoutes() []string {
	return cx.Mappings.ProxyRoutes
}

func (*configXml) GetEsFilterType(indexName string, typeName string) int {
	key, exist := cx.Mappings.FilterTypeMap[indexName+"-"+typeName]
	return util.If(exist == false, 0, key).(int)
}

func (*configXml) GetEsProjectId(indexName string, typeName string) int {
	key, exist := cx.Mappings.ProjectMap[indexName+"-"+typeName]
	return util.If(exist == false, 0, key).(int)
}

func (*configXml) loadXmlFile() {
	content, err := ioutil.ReadFile(setting.App.RootPath + "/config/config.xml")
	if err != nil {
		panic(err)
	}
	err = xml.Unmarshal(content, &cx)
	if err != nil {
		panic(err)
	}
	var key string
	for _, rc := range cx.Routes.Route {
		key = rc.IndexName + "-" + rc.TypeName
		_, isExist := cx.Mappings.RouteMap[key]
		if isExist {
			panic("配置存在相同的 key 请核查否则无法启动" + key)
		}
		cx.Mappings.RouteMap[key] = rc.DestinationIP + ":" + rc.DestinationPort + "/" + rc.DestinationIndexName + "/" + rc.DestinationTypeName + "/_search"
		cx.Mappings.ProxyRoutes = append(cx.Mappings.ProxyRoutes, rc.DestinationIndexName+"/"+rc.DestinationTypeName+"/_search")
		cx.Mappings.FilterTypeMap[key] = rc.FilterType
		cx.Mappings.ProjectMap[key] = rc.ProjectID
	}
}
