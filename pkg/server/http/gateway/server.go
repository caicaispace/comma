package admin

import (
	httpServer "comma/pkg/library/net/http"
	"comma/pkg/library/util/config"
	gatewayService "comma/pkg/service/gateway"
)

func NewServer(s *httpServer.Service) {
	routers := config.GetInstance().GetProxyRoutes()
	for _, route := range routers {
		s.Engine.Any(route, gatewayService.GetInstance().Dispatch)
	}
}
