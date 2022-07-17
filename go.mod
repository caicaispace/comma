module comma

go 1.16

replace golang.org/x/sys => github.com/golang/sys v0.0.0-20190322080309-f49334f85ddc

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/astaxie/beego v1.12.3
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/elastic/go-elasticsearch/v5 v5.6.1
	github.com/gin-gonic/gin v1.7.7
	github.com/go-playground/validator/v10 v10.10.0 // indirect
	github.com/go-resty/resty/v2 v2.7.0
	github.com/goinggo/mapstructure v0.0.0-20140717182941-194205d9b4a9
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kataras/tablewriter v0.0.0-20180708051242-e063d29b7c23
	github.com/landoop/tableprinter v0.0.0-20201125135848-89e81fc956e7
	github.com/lxmgo/config v0.0.0-20180313024057-8db99aca0f7e
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/mozillazg/go-pinyin v0.19.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/openzipkin-contrib/zipkin-go-opentracing v0.4.5
	github.com/openzipkin/zipkin-go v0.4.0
	github.com/penglongli/gin-metrics v0.1.9
	github.com/prometheus/client_golang v1.7.1
	github.com/prometheus/common v0.10.0
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	github.com/stretchr/testify v1.7.0
	github.com/ugorji/go v1.2.6 // indirect
	github.com/unknwon/com v1.0.1
	golang.org/x/crypto v0.0.0-20211215153901-e495a2d5b3d3 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	gopkg.in/yaml.v2 v2.4.0
	gorm.io/driver/mysql v1.2.2
	gorm.io/gorm v1.22.4
)
