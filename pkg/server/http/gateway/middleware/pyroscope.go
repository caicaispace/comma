package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/smallnest/zhuge"
)

type pyroscopeMiddleware struct{}

var profiler *zhuge.SampleProfiler

func NewPyroscope() *pyroscopeMiddleware {
	return &pyroscopeMiddleware{}
}

func (t pyroscopeMiddleware) Use(r *gin.Engine) {
	sc := zhuge.Config{
		ApplicationName: "comma.gateway",
		ServerAddress:   "http://127.0.0.1:4040",
		Tags:            map[string]string{"host": "127.0.0.1"},
		Logger:          zhuge.StandardLogger,
		AuthToken:       os.Getenv("PYROSCOPE_AUTH_TOKEN"),
		ProfileTypes: []zhuge.ProfileType{
			zhuge.ProfileCPU,
			zhuge.ProfileAllocObjects,
			zhuge.ProfileAllocSpace,
			zhuge.ProfileInuseObjects,
			zhuge.ProfileInuseSpace,
		},
	}
	var err error
	profiler, err = zhuge.NewSampleProfiler(sc)
	if err != nil {
		panic(err)
	}
	r.Use(t.handle)
}

func (t *pyroscopeMiddleware) handle(context *gin.Context) {
	context.Next()
}
