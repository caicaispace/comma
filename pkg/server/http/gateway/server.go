package admin

import (
	gatewayService "comma/pkg/service/gateway"

	"github.com/caicaispace/gohelper/config"
	httpServer "github.com/caicaispace/gohelper/server/http"
)

func NewServer(s *httpServer.Server) {
	routers := config.GetInstance().GetProxyRoutes()
	for _, route := range routers {
		s.Engine.Any(route, gatewayService.GetInstance().Dispatch)
	}
}
