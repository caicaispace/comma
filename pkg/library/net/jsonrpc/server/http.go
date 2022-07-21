package server

import (
	"comma/pkg/library/net/jsonrpc/common"
	"comma/pkg/library/net/jsonrpc/components/rate_limit"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type Http struct {
	Ip                    string
	Port                  string
	Server                common.Server
	ServerBeforeStartFunc func()
	ServerAfterStartFunc  func()
	Options               HttpOptions
	HttpRoutersMap        *map[string]func(w http.ResponseWriter, r *http.Request)
}

type HttpOptions struct {
	RateLimit    float64
	RateLimitMax int64
}

func NewHttpServer(ip string, port string) *Http {
	options := HttpOptions{
		0,
		0,
	}
	httpRouters := make(map[string]func(w http.ResponseWriter, r *http.Request))
	rateLimit := &rate_limit.RateLimit{}
	return &Http{
		Ip:   ip,
		Port: port,
		Server: common.Server{
			Sm:        sync.Map{},
			Hooks:     common.Hooks{},
			RateLimit: rateLimit,
		},
		Options:        options,
		HttpRoutersMap: &httpRouters,
	}
}

func (p *Http) Start() {
	if p.ServerBeforeStartFunc != nil {
		p.ServerBeforeStartFunc()
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", p.handleFunc)
	for k, v := range *p.HttpRoutersMap {
		mux.HandleFunc(k, v)
	}
	url := fmt.Sprintf("%s:%s", p.Ip, p.Port)
	// log.Printf("Listening http://%s:%s", p.Ip, p.Port)
	if p.ServerAfterStartFunc != nil {
		p.ServerAfterStartFunc()
	}
	http.ListenAndServe(url, mux)
}

func (p *Http) SetHttpRouters(routersMap *map[string]func(w http.ResponseWriter, r *http.Request)) {
	p.HttpRoutersMap = routersMap
}

func (p *Http) Register(s interface{}) {
	p.Server.Register(s)
}

func (p *Http) SetOptions(httpOptions interface{}) {
	p.Options = httpOptions.(HttpOptions)
}

func (p *Http) SetRateLimit(rate float64, max int64) {
	p.Server.RateLimit.GetBucket(rate, max)
}

func (p *Http) SetServerBeforeStartFunc(beforeFunc func()) {
	p.ServerBeforeStartFunc = beforeFunc
}

func (p *Http) SetServerAfterStartFunc(afterFunc func()) {
	p.ServerAfterStartFunc = afterFunc
}

func (p *Http) SetBeforeFunc(beforeFunc func(id interface{}, method string, params interface{}) error) {
	p.Server.Hooks.BeforeFunc = beforeFunc
}

func (p *Http) SetAfterFunc(afterFunc func(id interface{}, method string, result interface{}) error) {
	p.Server.Hooks.AfterFunc = afterFunc
}

func (p *Http) handleFunc(w http.ResponseWriter, r *http.Request) {
	var (
		err  error
		data []byte
	)
	w.Header().Set("Content-Type", "application/json")
	if data, err = ioutil.ReadAll(r.Body); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	res := p.Server.Handler(data)
	w.Write(res)
}
