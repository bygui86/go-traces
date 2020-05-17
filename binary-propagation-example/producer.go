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
	tracer, closer := Init("producer")
	defer closer.Close()

	span := tracer.StartSpan("producer")
	span.SetTag("Kind", "producer")

	var ctxBuffer bytes.Buffer
	fmt.Printf("%+v\n", span.Context())
	err := tracer.Inject(span.Context(), opentracing.Binary, &ctxBuffer)
	if err != nil {
		log.Fatal("Cannot inject span context", err.Error())
	}
	fmt.Println(ctxBuffer.String())
	fmt.Println(ctxBuffer.Bytes())

	err = ioutil.WriteFile("../spanctx", ctxBuffer.Bytes(), 0666)
	if err != nil {
		log.Fatal("Cannot store span context to file")
	}
	fmt.Println("I am a producer")
	span.LogKV("event", "producer-hello", "message", "inject-to-binary")
}
