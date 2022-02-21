package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/opentracing/opentracing-go/ext"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

var Trace opentracing.Tracer
var Closer io.Closer

func InitJaeger() (opentracing.Tracer, io.Closer) {

	c := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		}, // SamplingServerURL: "http://localhost:5778/sampling"

		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 60 * time.Second,
			LocalAgentHostPort:  "localhost:6831",
		},
		ServiceName: "test-aaaa",
	}
	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	tracer, closer, er := c.NewTracer(
		config.Logger(jLogger),
		config.Metrics(jMetricsFactory),
	)
	fmt.Println(33)
	fmt.Println(er)
	opentracing.SetGlobalTracer(tracer)
	//defer closer.Close()
	/*	tracer, closer, err := c.NewTracer(config.Logger(jaeger.StdLogger))
		if err != nil {
			panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
		}*/

	return tracer, closer
}

func main() {
	InitJaeger()
	tracer := opentracing.GlobalTracer()
	fmt.Println(111)
	fmt.Println(tracer)
	clientSpan := tracer.StartSpan("client")
	defer clientSpan.Finish()

	url := "http://baidu.com"
	req, _ := http.NewRequest("GET", url, nil)

	// Set some tags on the clientSpan to annotate that it's the client span. The additional HTTP tags are useful for debugging purposes.
	ext.SpanKindRPCClient.Set(clientSpan)
	ext.HTTPUrl.Set(clientSpan, url)
	ext.HTTPMethod.Set(clientSpan, "GET")

	// Inject the client span context into the headers
	tracer.Inject(clientSpan.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))
	resp, er := http.DefaultClient.Do(req)

	fmt.Println(er)
	fmt.Println(resp)
}
