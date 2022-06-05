package gateway

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"

	"goaway/pkg/library/core/l"
	httpServer "goaway/pkg/library/net/http"
	"goaway/pkg/library/util"
	"goaway/pkg/library/util/config"
	"goaway/pkg/service/segment"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

type Service struct {
	Addr      string
	AuthStr   string
	Segmenter *segment.SegmenterService
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
		}
		filterService = &FilterService{
			Segmenter: task.Segmenter,
		}
	})
	return service
}

// DictWithHttp 字典
func (ps *Service) DictWithHttp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ok bool
	var words []string
	mode := 0
	if modes, ok := r.URL.Query()["mode"]; ok {
		mode, _ = strconv.Atoi(modes[0])
	}
	words, ok = r.URL.Query()["word"]
	if !ok {
		_, _ = w.Write([]byte(`{"status":1,"error":"word lost"}`))
		return
	}
	word, err := url.QueryUnescape(words[0])
	if err != nil {
		_, _ = w.Write([]byte(`{"status":1,"error":"word lost"}`))
		return
	}
	rspMap := make(map[string]interface{})
	rspMap["status"] = 0
	var body []byte
	switch mode {
	case 0:
		rspMap["result"] = ps.Segmenter.SegmentSearchMode(word)
		body, _ = json.Marshal(rspMap)
	case 1:
		wordMap, synMap, hypMap, _ := ps.Segmenter.SegmentIndexMode(word, false, 0, false)
		rspMap["word"] = *wordMap
		if len(*synMap) > 0 {
			rspMap["syn"] = *synMap
		}
		if len(*hypMap) > 0 {
			rspMap["hyp"] = *hypMap
		}
		rspMap["result"] = rspMap
		body, _ = json.Marshal(rspMap)
	default:
		body = []byte(`{"status":2,"error":"no action"}`)
	}
	_, _ = w.Write(body)
}

func (ps *Service) DispatchWithRpc(esIndex, esType, esBody, handleType string) ([]byte, error) {
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
		l.Infof("url:%s cost:%dms param:%s", esRequestUrl, util.NowTimestampMS()-startTimeMS, esBody)
	case PASS_SEARCH_TYPE:
		l.Infof("%s 透传 url:%s cost:%dms ", handleType, esRequestUrl, util.NowTimestampMS()-startTimeMS)
	}
	return body, nil
}

func (ps *Service) DispatchWithGin(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			l.Infof("http request %v", err)
		}
	}()
	ctx := httpServer.Context{C: c}
	startTimeMS := util.NowTimestampMS()
	esRequestObj := ps.getEsRequestObjWithHttpServer(c.Request)
	if esRequestObj.Banned != "" {
		ctx.C.JSON(http.StatusOK, nil)
		return
	}
	if esRequestObj.SearchType == PROXY_SEARCH_TYPE {
		esRequestUrl := config.GetInstance().GetEsRoute(esRequestObj.IndexName, esRequestObj.TypeName)
		if esRequestUrl == "" {
			l.Errorf("%s not map url:%s cost %d ms param %s", c.Request.Method, c.Request.URL.Path, util.NowTimestampMS()-startTimeMS, esRequestObj.BodyRaw)
			ctx.C.JSON(http.StatusOK, nil)
			return
		}
		body, err := ps.reqElasticSearch(esRequestObj)
		if err != nil {
			l.Errorf("http request %v", err)
			ctx.C.JSON(http.StatusOK, body)
			return
		}
		l.Infof("url:%s cost:%dms param:%s", c.Request.URL.Path, util.NowTimestampMS()-startTimeMS, esRequestObj.BodyRaw)
		ctx.C.JSON(http.StatusOK, body)
	} else {
		body, err := ps.reqElasticSearch(esRequestObj)
		if err != nil {
			l.Errorf("http request %v", err)
		}
		l.Infof("%s pass url:%s cost:%dms ", c.Request.Method, c.Request.URL.Path, util.NowTimestampMS()-startTimeMS)
		ctx.C.JSON(http.StatusOK, body)
	}
}

func (ps *Service) DispatchWithHttp(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			l.Infof("http request %v", err)
		}
	}()
	r.URL.Path = strings.Replace(r.URL.Path, "/gateway", "", 1)
	startTimeMS := util.NowTimestampMS()
	esRequestObj := ps.getEsRequestObjWithHttpServer(r)
	body := []byte("")
	w.Header().Set("Content-Type", "application/json")
	if esRequestObj.Banned != "" {
		_, _ = w.Write(body)
		return
	}
	if esRequestObj.SearchType == PROXY_SEARCH_TYPE {
		esRequestUrl := config.GetInstance().GetEsRoute(esRequestObj.IndexName, esRequestObj.TypeName)
		if esRequestUrl == "" {
			l.Errorf("%s not map url:%s cost %d ms param %s", r.Method, r.URL.Path, util.NowTimestampMS()-startTimeMS, esRequestObj.BodyRaw)
			_, _ = w.Write(body)
			return
		}
		body, err := ps.reqElasticSearch(esRequestObj)
		if err != nil {
			l.Errorf("http request %v", err)
			_, _ = w.Write(body)
			return
		}
		l.Infof("url:%s cost:%dms param:%s", r.URL.Path, util.NowTimestampMS()-startTimeMS, esRequestObj.BodyRaw)
		_, _ = w.Write(body)
	} else {
		body, err := ps.reqElasticSearch(esRequestObj)
		if err != nil {
			l.Errorf("http request %v", err)
		}
		l.Infof("%s pass url:%s cost:%dms ", r.Method, r.URL.Path, util.NowTimestampMS()-startTimeMS)
		_, _ = w.Write(body)
	}
}

// 获取 es 请求数据对象
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

// 向 es 发送 http 请求
func (ps *Service) reqElasticSearchV2(esRequestObj *EsRequest) ([]byte, error) {
	fmt.Println("---------------------------------- " + util.NowDateTime() + " ----------------------------------")
	fmt.Println(esRequestObj.BodyRaw)
	// Create a Resty Client
	client := resty.New()
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

// 向 es 发送 http 请求
func (ps *Service) reqElasticSearch(esRequestObj *EsRequest) ([]byte, error) {
	fmt.Println("---------------------------------- " + util.NowDateTime() + " ----------------------------------")
	fmt.Println(esRequestObj.BodyRaw)
	// return []byte(""), nil
	var errStr string
	esHttpRequest, err := http.NewRequest(esRequestObj.RequestMethod, esRequestObj.RequestUrl, strings.NewReader(esRequestObj.BodyRaw))
	if err != nil {
		errStr = fmt.Sprintf("url:%s NewRequest %v ", esRequestObj.RequestUrl, err)
		l.Error(errStr)
		return nil, errors.New(errStr)
	}
	esHttpRequest.Header.Set("Content-Type", "application/json")
	esHttpRequest.Header.Set("Authorization", "Basic "+ps.AuthStr)
	// esHttpRequest.Header.Add("Connection", "close") // 等效的关闭方式
	// esHttpRequest.Close = true
	httpClient := &http.Client{}
	esRsp, err := httpClient.Do(esHttpRequest)
	if err != nil {
		errStr = fmt.Sprintf("url:%s Do %v", esRequestObj.RequestUrl, err)
		l.Error(errStr)
		return nil, errors.New(errStr)
	}
	defer esRsp.Body.Close()
	body, err := ioutil.ReadAll(esRsp.Body)
	if err != nil {
		errStr = fmt.Sprintf("url:%s httpClient.Do ReadAll error %s ", esRequestObj.RequestUrl, err)
		l.Error(errStr)
		return nil, errors.New(errStr)
	}
	return body, nil
}

func (ps *Service) getEsRequestUrl(reqUri string) string {
	return ps.Addr + reqUri
}
