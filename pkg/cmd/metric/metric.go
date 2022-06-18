package main

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	typeRequestAll     = "all"
	typeRequestFail    = "fail"
	typeRequestSucceed = "succeed"
	typeRequestLimit   = "limit"
	typeRequestReject  = "reject"
)

var (
	apiRequestCounterVec = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "goaway",
			Subsystem: "goaway",
			Name:      "api_request_total",
			Help:      "Total number of request made.",
		}, []string{"name", "type"})

	apiResponseHistogramVec = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "goaway",
			Subsystem: "goaway",
			Name:      "api_response_duration_seconds",
			Help:      "Bucketed histogram of api response time duration",
			Buckets:   prometheus.ExponentialBuckets(0.0005, 2.0, 20),
		}, []string{"name"})
)

type Metric struct{}

func init() {
	prometheus.Register(apiRequestCounterVec)
	prometheus.Register(apiResponseHistogramVec)
}

func (m *Metric) PostRequest(api string, startAt time.Time) {
	doMetrics := true
	incrRequestReject("test")
	if doMetrics {
		incrRequestSucceed("test")
		observeAPIResponse("test", startAt)
	}
}

func incrRequest(name string) {
	apiRequestCounterVec.WithLabelValues(name, typeRequestAll).Inc()
}

func incrRequestFailed(name string) {
	apiRequestCounterVec.WithLabelValues(name, typeRequestFail).Inc()
}

func incrRequestSucceed(name string) {
	apiRequestCounterVec.WithLabelValues(name, typeRequestSucceed).Inc()
}

func incrRequestReject(name string) {
	apiRequestCounterVec.WithLabelValues(name, typeRequestReject).Inc()
}

func observeAPIResponse(name string, startAt time.Time) {
	now := time.Now()
	apiResponseHistogramVec.WithLabelValues(name).Observe(now.Sub(startAt).Seconds())
}
