package admin

import (
	httpServer "goaway/pkg/library/net/http"
	"goaway/pkg/library/util/config"
	gatewayService "goaway/pkg/service/gateway"
)

func NewServer(s *httpServer.Service) {
	routers := config.GetInstance().GetProxyRoutes()
	for _, route := range routers {
		s.Engine.Any(route, gatewayService.GetInstance().Dispatch)
	}
}
