package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// 	"os/signal"
// 	"syscall"
// 	"time"
// )

// func main() {
// 	s := http.NewServeMux()
// 	s.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println(r.URL)
// 		time.Sleep(5 * time.Second)
// 		w.Write([]byte("Hello world!"))
// 	})
// 	server := &http.Server{
// 		Addr:    "127.0.0.1:8090",
// 		Handler: s,
// 	}
// 	go server.ListenAndServe()

// 	listenSignal(context.Background(), server)
// }

// func listenSignal(ctx context.Context, httpSrv *http.Server) {
// 	sigs := make(chan os.Signal, 1)
// 	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
// 	select {
// 	case <-sigs:
// 		log.Println("notify sigs")
// 		httpSrv.Shutdown(ctx)
// 		// 打印进程id
// 		log.Println("PID:", os.Getppid())
// 		log.Println("http shutdown")
// 	}
// }

// func everLoop(ctx context.Context) {
// LOOP:
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			// 收到信号退出无限循环
// 			break LOOP
// 		default:
// 			// 用一个 sleep 模拟业务逻辑
// 			time.Sleep(time.Second * 10)
// 		}
// 	}
// }

// func main() {
// 	// 监控系统信号和创建 Context 现在一步搞定
// 	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
// 	// 在收到信号的时候，会自动触发 ctx 的 Done ，这个 stop 是不再捕获注册的信号的意思，算是一种释放资源。
// 	defer stop()

// 	// 开始无限循环，收到信号就会退出
// 	everLoop(ctx)
// 	fmt.Println("graceful shuwdown")
// }
