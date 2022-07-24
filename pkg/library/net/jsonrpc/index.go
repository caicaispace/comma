package jsonrpc

import (
	"net/http"
	"strings"

	"comma/pkg/library/setting"
	"comma/pkg/library/util"

	util2 "comma/pkg/library/util"
)

type service struct {
	server                ServerInterface
	serverAddr            string
	serverProtocol        string
	serverBeforeStartFunc func()
	serverAfterStartFunc  func()
	serverBeforeFunc      func(id interface{}, method string, params interface{}) error
	serverAfterFunc       func(id interface{}, method string, params interface{}) error
	services              []interface{}
	routers               *map[string]func(w http.ResponseWriter, r *http.Request)
}

func NewServer() *service {
	s := &service{
		services: make([]interface{}, 0),
	}
	return s
}

func (s *service) RegisterService(service interface{}) *service {
	s.services = append(s.services, service)
	return s
}

func (s *service) RegisterRouters(routers *map[string]func(w http.ResponseWriter, r *http.Request)) *service {
	s.routers = routers
	return s
}

func (s *service) SetServerAfterStartFunc(afterFunc func()) {
	s.serverAfterStartFunc = afterFunc
}

func (s *service) SetServerBeforeStartFunc(afterFunc func()) {
	s.serverBeforeStartFunc = afterFunc
}

// Set the hook function of before method execution
func (s *service) SetBeforeFunc(fn func(id interface{}, method string, params interface{}) error) *service {
	// If the function returns an error, the program stops execution and returns an error message to the client
	// return errors.New("Custom Error")
	s.serverBeforeFunc = fn
	return s
}

// Set the hook function of after method execution
func (s *service) SetAfterFunc(fn func(id interface{}, method string, params interface{}) error) *service {
	// If the function returns an error, the program stops execution and returns an error message to the client
	// return errors.New("Custom Error")
	s.serverAfterFunc = fn
	return s
}

func (s *service) SetServerAddr(addr string) *service {
	s.serverAddr = addr
	return s
}

func (s *service) Start() {
	if s.serverAddr == "" {
		s.serverAddr = "127.0.0.1:8081"
	}
	if s.serverProtocol == "" {
		s.serverProtocol = "http"
	}
	util.CommandPrint(util.CommandSetPrintData("jsonrpc: "+s.serverProtocol, s.serverAddr, setting.Server.RunMode))
	host := strings.Split(s.serverAddr, ":")[0]
	port := strings.Split(s.serverAddr, ":")[1]
	js, err := New(
		s.serverProtocol,
		util2.If(host != "", host, setting.Server.Host).(string),
		util2.If(port != "", port, setting.Server.Port).(string),
	)
	if err != nil {
		panic(err)
	}
	s.server = js
	for _, service := range s.services {
		s.server.Register(service)
	}
	if s.serverProtocol == "http" {
		s.server.SetHttpRouters(httpHealthyCheck())
	}
	if s.routers != nil {
		s.server.SetHttpRouters(s.routers)
	}
	if s.serverBeforeFunc != nil {
		s.server.SetBeforeFunc(s.serverBeforeFunc)
	}
	if s.serverAfterFunc != nil {
		s.server.SetAfterFunc(s.serverAfterFunc)
	}
	if s.serverBeforeStartFunc != nil {
		s.server.SetServerBeforeStartFunc(s.serverBeforeStartFunc)
	}
	if s.serverAfterStartFunc != nil {
		s.server.SetServerAfterStartFunc(s.serverAfterStartFunc)
	}
	// s, _ := jsonrpcServer.NewServer("tcp", "127.0.0.1", "3232") // the protocol is tcp
	// s.SetOptions(js.TcpOptions{"aaaaaa", 2 * 1024 * 1024}) // Custom package EOF when the protocol is tcp
	// s.SetRateLimit(20, 10) //The maximum concurrent number is 10, The maximum request speed is 20 times per second
	s.server.Start()
}

func httpHealthyCheck() *map[string]func(w http.ResponseWriter, r *http.Request) {
	return &map[string]func(w http.ResponseWriter, r *http.Request){
		"/ping": func(w http.ResponseWriter, r *http.Request) {
			body := []byte("pong")
			w.Write(body)
		},
		"/check": func(w http.ResponseWriter, r *http.Request) {
			body := []byte("ok")
			w.Write(body)
		},
		// "/json": func(w http.ResponseWriter, r *http.Request) {
		// 	jsonBody, _ := json.Marshal(map[string]interface{}{
		// 		"url":       r.URL.String(),
		// 		"paramsStr": r.URL.RawQuery,
		// 		"params":    r.URL.Query(),
		// 		"method":    r.Method,
		// 		"header":    r.Body,
		// 	})
		// 	w.Header().Set("Content-Type", "application/json")
		// 	w.Write(jsonBody)
		// },
	}
}
