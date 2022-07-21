package http

import (
	"comma/pkg/library/core/l"
	"comma/pkg/library/net/http/middleware"
	"comma/pkg/library/setting"
	"comma/pkg/library/util"
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	// server
	serverMode = flag.String("mode", gin.DebugMode, "Server: run mode")
	count      int64
)

func init() {
	setting.Server.RunMode = *serverMode
}

type Service struct {
	Engine     *gin.Engine
	ServerAddr string
	beforeFunc func(env string)
}

func (s *Service) registerDefaultRouter() {
	s.Engine.GET("/check", func(c *gin.Context) {
		c.String(http.StatusOK, "ok "+fmt.Sprint(count)+" remote:"+c.Request.RemoteAddr+" "+c.Request.URL.String())
		count++
	})
	s.Engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
}

func NewServer() *Service {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard // disable router map log
	return &Service{
		Engine: gin.New(),
		// Engine: gin.Default(),
	}
}

func (s *Service) SetBeforeFunc(fn func(env string)) *Service {
	s.beforeFunc = fn
	return s
}

func (s *Service) SetRouters() {
}

func (s *Service) SetServerAddr(addr string) *Service {
	s.ServerAddr = addr
	return s
}

// zipkin
func (s *Service) UseTrace(zipkinAddr, serviceName, serviceAddr string) *Service {
	trace := middleware.NewTraceV2(zipkinAddr, serviceName, serviceAddr)
	s.Engine.Use(func(c *gin.Context) {
		span := (trace.ZipkinTracer).StartSpan(c.FullPath())
		defer span.Finish()
		c.Next()
	})
	// defer (trace.ZipkinReporter).Close()
	return s
}

func (s *Service) UseGrafana() {
	middleware.NewGrafana(s.Engine)
}

// var f *os.File

func listenSignal(ctx context.Context, httpSrv *http.Server) {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	sig := <-sc
	l.Infof("exit: signal=<%d>.", sig)
	switch sig {
	case syscall.SIGTERM:
		l.Infof("exit: bye :-).")
		os.Exit(0)
	default:
		l.Infof("exit: bye :-(.")
		// // CPU 性能分析
		// f.Close()
		// pprof.StopCPUProfile()
		os.Exit(1)
	}
}

func (s *Service) Start() {
	// //CPU 性能分析
	// runtime.GOMAXPROCS(1)              // 限制 CPU 使用数，避免过载
	// runtime.SetMutexProfileFraction(1) // 开启对锁调用的跟踪
	// runtime.SetBlockProfileRate(1)     // 开启对阻塞操作的跟踪
	// f, err := os.OpenFile("cpu.prof", os.O_RDWR|os.O_CREATE, 0644)
	// if err != nil {
	// 	l.Error(err)
	// 	return
	// }
	// pprof.StartCPUProfile(f)
	s.Engine.Use(gin.Logger())
	s.Engine.Use(gin.Recovery())
	if s.beforeFunc != nil {
		s.beforeFunc(setting.Server.Env)
	}
	if s.ServerAddr != "" {
		setting.Server.Addr = s.ServerAddr
	}
	s.registerDefaultRouter()
	util.CommandPrint(util.CommandSetPrintData("restful", setting.Server.Addr, setting.Server.RunMode))
	maxHeaderBytes := 1 << 20
	gin.SetMode(setting.Server.RunMode)
	httpServer := &http.Server{
		// Addr: "0.0.0.0:" + strings.Split(setting.Server.Addr, ":")[1],
		Addr:           setting.Server.Addr,
		Handler:        s.Engine,
		ReadTimeout:    time.Duration(setting.Server.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(setting.Server.WriteTimeout) * time.Second,
		MaxHeaderBytes: maxHeaderBytes,
	}
	go httpServer.ListenAndServe()
	listenSignal(context.Background(), httpServer)
}
