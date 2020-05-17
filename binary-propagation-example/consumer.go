package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"

	"github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
	config "github.com/uber/jaeger-client-go/config"
)

func Init(service string) (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}
	tracer, closer, err := cfg.New(service, config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	return tracer, closer
}
func main() {
	tracer, closer := Init("consumer")
	defer closer.Close()
	spanBytes, err := ioutil.ReadFile("../spanctx")
	if err != nil {
		log.Fatal("Cannot read span context file")
	}
	fmt.Println(string(spanBytes))
	fmt.Println(spanBytes)
	var ctxBuffer bytes.Buffer
	ctxBuffer.Read(spanBytes)
	spanCtx, err := tracer.Extract(opentracing.Binary, &ctxBuffer)
	if err != nil {
		log.Fatal("Cannot extract span context", err.Error())
	}
	span := tracer.StartSpan("consumer", opentracing.FollowsFrom(spanCtx))
	span.SetTag("Kind", "consumer")
	span.LogKV("event", "consumer-hello", "message", "extraced-from-binary")
	fmt.Println("I am consumer")
}
