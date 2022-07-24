package config

import (
	"fmt"
	"os"
	"sync"

	"comma/pkg/library/setting"
	"comma/pkg/library/util"

	"github.com/BurntSushi/toml"
)

type Conf struct {
	Title  string
	Desc   string
	Env    string
	Server server
	DB     db
	Metric metric
	ES     es
}

type server struct {
	Host string
}

type db struct {
	Addr string
}

type metric struct {
	Enable       bool
	Job          string
	Instance     string
	Address      string
	IntervalSync uint64
}

type es struct {
	Addr          string
	Username      string
	Password      string
	Routers       []esRouters
	proxyRouters  []string
	routerMap     map[string]string
	filterTypeMap map[string]uint8
	projectMap    map[string]uint8
}

type esRouters struct {
	Addr       string
	Index      string
	Type       string
	FilterType uint8
	ProjectId  uint8
}

var (
	conf     *Conf
	confOnce sync.Once
)

func GetInstance() *Conf {
	confOnce.Do(func() {
		conf = &Conf{
			ES: es{
				Routers:       make([]esRouters, 0),
				proxyRouters:  make([]string, 0),
				routerMap:     make(map[string]string),
				filterTypeMap: make(map[string]uint8),
				projectMap:    make(map[string]uint8),
			},
		}
		conf.loadFile()
	})
	return conf
}

func (c *Conf) GetEnv() string {
	return c.Env
}

func (c *Conf) GetServerHost() string {
	return c.Server.Host
	// localIp := setting.Server.Host
	// if runtime.GOOS == "linux" {
	// 	localIp = util.LocalIP()
	// }
	// return localIp
}

func (c *Conf) GetEs() *es {
	return &c.ES
}

func (c *Conf) GetMetricIsEnable() bool {
	return c.Metric.Enable
}

func (c *Conf) GetMetric() *metric {
	return &c.Metric
}

func (c *Conf) GetProxyRoutes() []string {
	return c.ES.proxyRouters
}

func (c *Conf) GetEsRoute(indexName, typeName string) string {
	key, exist := c.ES.routerMap[indexName+"-"+typeName]
	return util.If(exist == false, "", key).(string)
}

func (c *Conf) GetEsFilterType(indexName, typeName string) uint8 {
	key, exist := c.ES.filterTypeMap[indexName+"-"+typeName]
	return util.If(exist == false, 0, key).(uint8)
}

func (c *Conf) GetEsProjectId(indexName, typeName string) uint8 {
	key, exist := c.ES.projectMap[indexName+"-"+typeName]
	return util.If(exist == false, 0, key).(uint8)
}

func (c *Conf) GetDB() string {
	return c.DB.Addr
}

func (c *Conf) loadFile() {
	var err error
	f := util.GetRootPath() + "/config/conf.toml"
	if _, err = os.Stat(f); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	_, err = toml.DecodeFile(f, conf)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if c.Env != "" {
		setting.Server.Env = c.Env
	}
	var key, path string
	for _, router := range conf.ES.Routers {
		key = router.Index + "-" + router.Type
		_, isExist := conf.ES.routerMap[key]
		if isExist {
			panic("please verify that the same key is configured or it will not start" + key)
		}
		path = router.Index + "/" + router.Type + "/_search"
		conf.ES.proxyRouters = append(conf.ES.proxyRouters, path)
		conf.ES.routerMap[key] = router.Addr + "/" + path
		conf.ES.filterTypeMap[key] = router.FilterType
		conf.ES.projectMap[key] = router.ProjectId
	}
}
