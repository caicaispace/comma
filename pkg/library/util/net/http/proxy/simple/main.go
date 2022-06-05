package main

import (
	"log"
	"net/http"

	"goaway/pkg/library/util/net/http/proxy"
)

func OnError(ctx *proxy.Context, where string,
	err *proxy.Error, opErr error,
) {
	// Log errors.
	log.Printf("ERR: %s: %s [%s]", where, err, opErr)
}

func OnAccept(ctx *proxy.Context, w http.ResponseWriter,
	r *http.Request,
) bool {
	// Handle local request has path "/info"
	if r.Method == "GET" && !r.URL.IsAbs() && r.URL.Path == "/info" {
		w.Write([]byte("This is go-httpproxy."))
		return true
	}
	return false
}

func OnAuth(ctx *proxy.Context, authType string, user string, pass string) bool {
	// Auth test user.
	if user == "test" && pass == "1234" {
		return true
	}
	return false
}

func OnConnect(ctx *proxy.Context, host string) (
	ConnectAction proxy.ConnectAction, newHost string,
) {
	// Apply "Man in the Middle" to all ssl connections. Never change host.
	return proxy.ConnectMitm, host
}

func OnRequest(ctx *proxy.Context, req *http.Request) (resp *http.Response) {
	// Log proxying requests.
	log.Printf("INFO: Proxy: %s %s", req.Method, req.URL.String())
	return
}

func OnResponse(ctx *proxy.Context, req *http.Request, resp *http.Response) {
	// Add header "Via: go-httpproxy".
	resp.Header.Add("Via", "go-httpproxy")
}

func main() {
	// Create a new proxy with default certificate pair.
	prx, _ := proxy.NewProxy()

	// Set handlers.
	prx.OnError = OnError
	prx.OnAccept = OnAccept
	prx.OnAuth = OnAuth
	prx.OnConnect = OnConnect
	prx.OnRequest = OnRequest
	prx.OnResponse = OnResponse

	// Listen...
	http.ListenAndServe("127.0.0.1:8080", prx)
}
