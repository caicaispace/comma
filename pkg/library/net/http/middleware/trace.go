package middleware

import (
	"log"

	"github.com/opentracing/opentracing-go"
	zipkinot "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	openzipkin "github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/reporter"
	zipkinHTTP "github.com/openzipkin/zipkin-go/reporter/http"
)

type Trace struct {
	ZipkinTracer   opentracing.Tracer
	ZipkinReporter reporter.Reporter
}

func NewTrace(zipkinAddr, serviceName, serviceAddr string) *Trace {
	zkReporter := zipkinHTTP.NewReporter(zipkinAddr)
	defer zkReporter.Close()
	endpoint, err := openzipkin.NewEndpoint(serviceName, serviceAddr)
	if err != nil {
		log.Panicf("unable to create local endpoint: %+v\n", err)
	}
	nativeTracer, err := zipkin.NewTracer(zkReporter, zipkin.WithLocalEndpoint(endpoint))
	if err != nil {
		log.Panicf("unable to create tracer: %+v\n", err)
	}
	zkTracer := zipkinot.Wrap(nativeTracer)
	opentracing.SetGlobalTracer(zkTracer)
	return &Trace{
		ZipkinTracer: zkTracer,
	}
}

func NewTraceV2(zipkinAddr, serviceName, serviceAddr string) *Trace {
	zkReporter := zipkinHTTP.NewReporter(zipkinAddr)
	endpoint, err := openzipkin.NewEndpoint(serviceName, serviceAddr)
	if err != nil {
		log.Fatalf("unable to create local endpoint: %+v\n", err)
		return nil
	}
	nativeTracer, err := openzipkin.NewTracer(zkReporter, openzipkin.WithTraceID128Bit(true), openzipkin.WithLocalEndpoint(endpoint))
	if err != nil {
		log.Fatalf("unable to create tracer: %+v\n", err)
		return nil
	}
	zkTracer := zipkinot.Wrap(nativeTracer)
	opentracing.SetGlobalTracer(zkTracer)
	return &Trace{
		ZipkinTracer:   zkTracer,
		ZipkinReporter: zkReporter,
	}
}
