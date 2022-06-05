package admin

import (
	"net/http"
	// httpproxy "goaway/pkg/library/util/net/http/proxy"

	gatewayService "goaway/pkg/service/gateway"
)

// ---------------------- http ----------------------

func NewServer(addr string) {
	http.HandleFunc("/", gatewayService.GetInstance().DispatchWithHttp)
	http.ListenAndServe(addr, nil)
}

// ---------------------- proxy ----------------------

// func OnError(ctx *httpproxy.Context, where string, err *httpproxy.Error, opErr error) {
// 	// Log errors.
// 	log.Printf("ERR: %s: %s [%s]", where, err, opErr)
// }

// func OnAccept(ctx *httpproxy.Context, w http.ResponseWriter, r *http.Request) bool {
// 	// Handle local request has path "/info"
// 	if r.Method == "GET" && !r.URL.IsAbs() && r.URL.Path == "/" {
// 		w.Write([]byte("This is go-httpproxy."))
// 		return true
// 	}
// 	return false
// }

// func NewServer() {
// 	// Create a new proxy with default certificate pair.
// 	prx, _ := httpproxy.NewProxy()
// 	// Set handlers.
// 	prx.OnError = OnError
// 	prx.OnAccept = OnAccept
// 	// Listen...
// 	http.ListenAndServe("127.0.0.1:8081", prx)
// }

// ---------------------- gin ----------------------

// func NewServer(s *httpServer.Service) {
// 	routers := config.GetXmlInstance().GetProxyRoutes()
// 	for _, route := range routers {
// 		// s.Engine.Any(route, gatewayService.GetInstance().DispatchWithGin)
// 		s.Engine.POST(route, gatewayService.GetInstance().DispatchWithGin)
// 	}
// }
