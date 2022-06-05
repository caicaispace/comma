package pool

import (
	"net/http"
	"time"
)

type PoolV2 struct {
	maxPoolSize  int
	cClient      chan *http.Client
	reqPerSecond int
	rateLimiter  <-chan time.Time
}

func NewV2(maxPoolSize int, reqPerSec int) *PoolV2 {
	var client chan *http.Client = nil
	if maxPoolSize > 0 {
		client = make(chan *http.Client, maxPoolSize)
		// client <- new(http.Client)
		for i := 0; i < maxPoolSize; i++ {
			client <- new(http.Client)
		}
	}

	var emitter <-chan time.Time = nil
	if reqPerSec > 0 {
		emitter = time.NewTicker(time.Second / time.Duration(reqPerSec)).C // x req/s == 1s/x req (inverse)
	}

	return &PoolV2{
		maxPoolSize:  maxPoolSize,
		cClient:      client,
		reqPerSecond: reqPerSec,
		rateLimiter:  emitter,
	}
}

// 获取一个连接
func (c *PoolV2) Get() *http.Client {
	if c.cClient != nil {
		return <-c.cClient
	}
	return nil
}

// 回收连接
func (c *PoolV2) Recycle(httpClient *http.Client) {
	if httpClient != nil {
		c.cClient <- httpClient
	}
}

// 销毁全部连接池
func (c *PoolV2) Destroy() {
	if c.cClient != nil {
		for i := 0; i < len(c.cClient); i++ {
			c.cClient <- nil
		}
	}
}

// 获取连接池中的连接数
func (c *PoolV2) Len() int {
	if c.cClient != nil {
		return len(c.cClient)
	}
	return 0
}
