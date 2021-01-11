package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegerCfg "github.com/uber/jaeger-client-go/config"
)

func main() {
	fmt.Println("Start standalone")

	closer := prepareTracer()

	go func() {
		for {
			span := opentracing.StartSpan("simple-operation-main")
			doWork(opentracing.ContextWithSpan(context.Background(), span))
			span.Finish()
		}
	}()

	fmt.Println("Standalone up and running")
	startSysCallChannel()

	fmt.Println("Termination signal received! Timeout 3 seconds")
	closeTracer(closer)
	time.Sleep(5 * time.Second)
}

func prepareTracer() io.Closer {
	cfg, cfgErr := jaegerCfg.FromEnv()
	if cfgErr != nil {
		fmt.Printf("ERROR - Get Jaeger configs failed: %s", cfgErr.Error())
		panic(cfgErr)
	}

	tracer, closer, tracerErr := cfg.NewTracer()
	if tracerErr != nil {
		fmt.Printf("ERROR - Create new tracer failed: %s", tracerErr.Error())
		panic(tracerErr)
	}
	opentracing.SetGlobalTracer(tracer)
	return closer
}

func doWork(ctx context.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx, "simple-operation-doWork")
	defer span.Finish()

	// sleep to simulate work
	time.Sleep(3 * time.Second)

	var tracingMsg string
	traceID, sampled := ExtractSampledTraceID(ctx)
	if !sampled {
		tracingMsg = "traceID not sampled"
	} else {
		tracingMsg = fmt.Sprintf("traceID %s", traceID)
	}
	fmt.Printf("INFO - Work done - %s \n", tracingMsg)
}

// ExtractSampledTraceID works like ExtractTraceID but the returned bool is only true if the returned trace id is sampled.
// copied from https://github.com/weaveworks/common/blob/master/middleware/http_tracing.go
func ExtractSampledTraceID(ctx context.Context) (string, bool) {
	sp := opentracing.SpanFromContext(ctx)
	if sp == nil {
		return "", false
	}

	spanCtx, ok := sp.Context().(jaeger.SpanContext)
	if !ok {
		return "", false
	}

	return spanCtx.TraceID().String(), spanCtx.IsSampled()
}

func startSysCallChannel() {
	syscallCh := make(chan os.Signal)
	signal.Notify(syscallCh, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	<-syscallCh
}

func closeTracer(closer io.Closer) {
	err := closer.Close()
	if err != nil {
		fmt.Printf("ERROR - Close tracer failed: %s", err.Error())
	}
}
