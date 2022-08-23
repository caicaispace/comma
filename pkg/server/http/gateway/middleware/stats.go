package middleware

import (
	"github.com/arl/statsviz"
	"github.com/gin-gonic/gin"
)

type statsMiddleware struct{}

func NewStats() *statsMiddleware {
	return &statsMiddleware{}
}

func (t statsMiddleware) Use(r *gin.Engine) {
	r.GET("/debug/stats/*filepath", func(context *gin.Context) {
		if context.Param("filepath") == "/ws" {
			statsviz.Ws(context.Writer, context.Request)
			return
		}
		statsviz.IndexAtRoot("/debug/stats").ServeHTTP(context.Writer, context.Request)
	})
	r.Use(t.handle)
}

func (t *statsMiddleware) handle(context *gin.Context) {
	context.Next()
}
