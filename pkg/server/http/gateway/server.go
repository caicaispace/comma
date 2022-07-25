package admin

import (
	"comma/pkg/library/util/config"
	gatewayService "comma/pkg/service/gateway"

	httpServer "github.com/caicaispace/gohelper/server/http"
)

func NewServer(s *httpServer.Service) {
	routers := config.GetInstance().GetProxyRoutes()
	for _, route := range routers {
		s.Engine.Any(route, gatewayService.GetInstance().Dispatch)
	}
}
