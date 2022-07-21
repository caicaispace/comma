package server

import (
	"comma/pkg/library/net/jsonrpc/common"
	"comma/pkg/library/net/jsonrpc/components/rate_limit"
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"
)

type Tcp struct {
	Ip                    string
	Port                  string
	Server                common.Server
	ServerBeforeStartFunc func()
	ServerAfterStartFunc  func()
	Options               TcpOptions
	HttpRoutersMap        *map[string]func(w http.ResponseWriter, r *http.Request)
}

type TcpOptions struct {
	PackageEof       string
	PackageMaxLength int32
}

func NewTcpServer(ip string, port string) *Tcp {
	options := TcpOptions{
		"\r\n",
		1024 * 1024 * 2,
	}
	httpRouters := make(map[string]func(w http.ResponseWriter, r *http.Request))
	rateLimit := &rate_limit.RateLimit{}
	return &Tcp{
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

func (p *Tcp) Start() {
	if p.ServerBeforeStartFunc != nil {
		p.ServerBeforeStartFunc()
	}
	addr := fmt.Sprintf("%s:%s", p.Ip, p.Port)
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr) // 解析tcp服务
	if err != nil {
		common.Debug(err.Error())
	}
	listener, _ := net.ListenTCP("tcp", tcpAddr)
	// log.Printf("Listening tcp://%s:%s", p.Ip, p.Port)
	if p.ServerAfterStartFunc != nil {
		p.ServerAfterStartFunc()
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			common.Debug(err.Error())
			continue
		}
		go p.handleFunc(ctx, conn)
	}
}

func (p *Tcp) SetHttpRouters(routersMap *map[string]func(w http.ResponseWriter, r *http.Request)) {
	// p.HttpRoutersMap = routersMap
}

func (p *Tcp) Register(s interface{}) {
	p.Server.Register(s)
}

func (p *Tcp) SetOptions(tcpOptions interface{}) {
	p.Options = tcpOptions.(TcpOptions)
}

func (p *Tcp) SetRateLimit(rate float64, max int64) {
	p.Server.RateLimit.GetBucket(rate, max)
}

func (p *Tcp) SetServerAfterStartFunc(afterFunc func()) {
	p.ServerAfterStartFunc = afterFunc
}

func (p *Tcp) SetServerBeforeStartFunc(afterFunc func()) {
	p.ServerBeforeStartFunc = afterFunc
}

func (p *Tcp) SetBeforeFunc(beforeFunc func(id interface{}, method string, params interface{}) error) {
	p.Server.Hooks.BeforeFunc = beforeFunc
}

func (p *Tcp) SetAfterFunc(afterFunc func(id interface{}, method string, result interface{}) error) {
	p.Server.Hooks.AfterFunc = afterFunc
}

func (p *Tcp) handleFunc(ctx context.Context, conn net.Conn) {
	defer conn.Close()
	select {
	case <-ctx.Done():
		return
	default:
		//	do nothing
	}
	for {
		buf := make([]byte, p.Options.PackageMaxLength)
		n, err := conn.Read(buf)
		if err != nil {
			if n == 0 {
				continue
			}
			common.Debug(err.Error())
			break
		}
		l := len([]byte(p.Options.PackageEof))
		res := p.Server.Handler(buf[:n-l])
		res = append(res, []byte(p.Options.PackageEof)...)
		conn.Write(res)
	}
}
