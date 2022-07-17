package metric

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
			Namespace: "comma",
			Subsystem: "comma",
			Name:      "req_total",
			Help:      "Total number of request made.",
		}, []string{"name", "type"})

	apiResponseHistogramVec = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "comma",
			Subsystem: "comma",
			Name:      "rsp_duration_seconds",
			Help:      "Bucketed histogram of response time duration",
			Buckets:   prometheus.ExponentialBuckets(0.0005, 2.0, 20),
		}, []string{"name"})
)

type Metric struct{}

func init() {
	prometheus.Register(apiRequestCounterVec)
	prometheus.Register(apiResponseHistogramVec)
}

func NewMetric() *Metric {
	return &Metric{}
}

func (m *Metric) PostRequest(name string, satus bool, startAt time.Time) {
	doMetrics := true
	incrRequest(name)
	incrRequestReject(name)
	if satus == false {
		incrRequestFailed(name)
	}
	if doMetrics {
		incrRequestSucceed(name)
		observeAPIResponse(name, startAt)
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
