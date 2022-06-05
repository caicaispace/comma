package main

import (
	"goaway/pkg/demo/http/controller"
	httpServer "goaway/pkg/library/net/http"
)

func init() {
}

func main() {
	// yiigo.LoadEnv(yiigo.WithEnvWatcher(func(event fsnotify.Event) {
	// 	fmt.Println(os.Getenv("ENV"))
	// }))
	s := httpServer.NewServer()
	apiV1 := s.Engine.Group("/v1/api")
	{
		apiV1.GET("/test", controller.Test)
		apiV1.GET("/test_pager", controller.TestPager)
		apiV1.GET("/db", controller.DB)
	}
	s.Start()
}
