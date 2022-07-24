package gateway

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"

	"comma/pkg/library/core/l"
	"comma/pkg/library/util"
	"comma/pkg/library/util/config"
	"comma/pkg/library/util/metric"
	"comma/pkg/service/segment"

	httpServer "comma/pkg/library/net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

type Service struct {
	Addr      string
	AuthStr   string
	Segmenter *segment.SegmenterService
	Metric    *metric.Metric
}

var filterService *FilterService

var (
	service *Service
	once    sync.Once
)

func GetInstance() *Service {
	once.Do(func() {
		esConfig := config.GetInstance().GetEs()
		task := segment.GetTaskServiceInstance()
		service = &Service{
			Addr:      esConfig.Addr,
			AuthStr:   base64.StdEncoding.EncodeToString([]byte(esConfig.Username + ":" + esConfig.Password)),
			Segmenter: task.Segmenter,
			Metric:    metric.NewMetric(),
		}
		filterService = &FilterService{
			Segmenter: task.Segmenter,
		}
	})
	return service
}

func (ps *Service) Dispatch(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			l.Infof("http request %v", err)
		}
	}()
	startAt := time.Now()
	ctx := httpServer.Context{C: c}
	startTimeMS := util.NowTimestampMS()
	esRequestObj := ps.getEsRequestObjWithHttpServer(c.Request)
	if esRequestObj.Banned != "" {
		ctx.C.Writer.Write(nil)
		return
	}
	if esRequestObj.SearchType == PROXY_SEARCH_TYPE {
		esRequestUrl := config.GetInstance().GetEsRoute(esRequestObj.IndexName, esRequestObj.TypeName)
		if esRequestUrl == "" {
			l.Errorf("%s not map url:%s time:%d ms body %s", c.Request.Method, c.Request.URL.Path, util.NowTimestampMS()-startTimeMS, esRequestObj.BodyRaw)
			ctx.JOSN(nil)
			return
		}
		body, err := ps.reqElasticSearch(esRequestObj)
		if err != nil {
			l.Errorf("http request %v", err)
			ctx.JOSN(body)
			return
		}
		l.Infof("url:%s time:%dms body:%s", c.Request.URL.Path, util.NowTimestampMS()-startTimeMS, esRequestObj.BodyRaw)
		ctx.JOSN(body)
	} else {
		body, err := ps.reqElasticSearch(esRequestObj)
		if err != nil {
			l.Errorf("http request %v", err)
		}
		l.Infof("%s pass url:%s time:%dms ", c.Request.Method, c.Request.URL.Path, util.NowTimestampMS()-startTimeMS)
		ctx.JOSN(body)
	}
	go func() {
		ps.Metric.PostRequest(c.Request.RequestURI, true, startAt)
	}()
}

func (ps *Service) DispatchWithJsonRpc(esIndex, esType, esBody, handleType string) ([]byte, error) {
	startTimeMS := util.NowTimestampMS()
	esRequestUrl := ps.getEsRequestUrl("")
	esRequestObj := &EsRequest{
		SearchType: PASS_SEARCH_TYPE,
		RequestUrl: esRequestUrl,
		IndexName:  esIndex,
		TypeName:   esType,
		BodyRaw:    esBody,
		ProjectId:  config.GetInstance().GetEsProjectId(esIndex, esType),
	}
	if filter := filterService.GetFilterInstance(esRequestObj.IndexName, esRequestObj.TypeName); filter != nil {
		filter.FilterRequest(esRequestObj)
	}
	esRoute := config.GetInstance().GetEsRoute(esRequestObj.IndexName, esRequestObj.TypeName)
	switch handleType {
	case "search":
		esRequestObj.RequestMethod = "POST"
		if esRoute != "" {
			esRequestObj.RequestUrl = esRoute
		}
		esRequestObj.SearchType = PROXY_SEARCH_TYPE
	case "update":
		esRequestObj.RequestMethod = "PUT"
		if esRoute != "" {
			esRequestObj.RequestUrl = esRoute
		}
		esRequestObj.SearchType = PROXY_SEARCH_TYPE
	case "delete":
		esRequestObj.RequestMethod = "DELETE"
		esRequestObj.SearchType = PASS_SEARCH_TYPE
	default:
		esRequestObj.SearchType = PASS_SEARCH_TYPE
	}
	body, err := ps.reqElasticSearch(esRequestObj)
	if err != nil {
		l.Errorf("http request %v", err)
		return nil, err
	}
	switch esRequestObj.SearchType {
	case PROXY_SEARCH_TYPE:
		l.Infof("url:%s time:%dms body:%s", esRequestUrl, util.NowTimestampMS()-startTimeMS, esBody)
	case PASS_SEARCH_TYPE:
		l.Infof("%s transparent url:%s time:%dms ", handleType, esRequestUrl, util.NowTimestampMS()-startTimeMS)
	}
	return body, nil
}

// get es request object
func (ps *Service) getEsRequestObjWithHttpServer(r *http.Request) *EsRequest {
	esRequestObj := &EsRequest{
		RequestUrl:    ps.getEsRequestUrl(r.URL.RequestURI()),
		SearchType:    PROXY_SEARCH_TYPE,
		RequestMethod: r.Method,
	}
	if !strings.Contains(r.URL.Path, "_search") {
		esRequestObj.SearchType = PASS_SEARCH_TYPE
		l.Infof("pass url: %s", r.URL.Path)
		return esRequestObj
	}
	urlPathSlice := strings.Split(r.URL.Path, "/")
	if urlPathSlice == nil || len(urlPathSlice) != 4 {
		esRequestObj.SearchType = PASS_SEARCH_TYPE
		l.Errorf("fatal error url: %s", r.URL.Path)
		return esRequestObj
	}
	esRequestObj.IndexName = urlPathSlice[1]
	esRequestObj.TypeName = urlPathSlice[2]
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic("read post request body fail")
	}
	esRequestObj.BodyRaw = string(body)
	esRequestObj.ProjectId = config.GetInstance().GetEsProjectId(esRequestObj.IndexName, esRequestObj.TypeName)
	if filter := filterService.GetFilterInstance(esRequestObj.IndexName, esRequestObj.TypeName); filter != nil {
		filter.FilterRequest(esRequestObj)
	}
	return esRequestObj
}

var client = resty.New()

// send http request to es
func (ps *Service) reqElasticSearch(esRequestObj *EsRequest) ([]byte, error) {
	if config.GetInstance().GetEnv() == "dev" {
		fmt.Println("---------------------------------- " + util.NowDateTime() + " ----------------------------------")
		fmt.Println(esRequestObj.BodyRaw)
	}
	// Create a Resty Client
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Basic "+ps.AuthStr).
		SetBody(strings.NewReader(esRequestObj.BodyRaw)).
		Post(esRequestObj.RequestUrl)
	if err != nil {
		errStr := fmt.Sprintf("url:%s NewRequest %v ", esRequestObj.RequestUrl, err)
		l.Error(errStr)
		return nil, errors.New(errStr)
	}
	return resp.Body(), nil
}

func (ps *Service) getEsRequestUrl(reqUri string) string {
	return ps.Addr + reqUri
}
