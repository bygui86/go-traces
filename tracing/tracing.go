package tracing

import (
	"io"
	"os"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"

	"github.com/bygui86/go-traces/logging"
)

/*
	By default, the client sends traces via UDP to the agent at localhost:6831.
	Use JAEGER_AGENT_HOST and JAEGER_AGENT_PORT to send UDP traces to a different host:port.

	If JAEGER_ENDPOINT is set, the client sends traces to the endpoint via HTTP, making the JAEGER_AGENT_HOST and
	JAEGER_AGENT_PORT unused.

	If JAEGER_ENDPOINT is secured, HTTP basic authentication can be performed by setting the JAEGER_USER and
	JAEGER_PASSWORD environment variables.
*/

func InitSample(serviceName string) io.Closer {
	// Sample configuration for testing. Use constant sampling to sample every trace
	// and enable LogSpan to log every span via configured Logger.
	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
		},
	}

	// Initialize tracer with a logger and a metrics factory
	closer, tracerErr := cfg.InitGlobalTracer(
		serviceName,
	)
	if tracerErr != nil {
		logging.SugaredLog.Errorf("Tracer creation failed: %s", tracerErr.Error())
		os.Exit(501)
	}

	logging.SugaredLog.Infof("Global tracer registered: %t", opentracing.IsGlobalTracerRegistered())

	return closer
}

func InitTestingTracing(serviceName string) io.Closer {
	// Sample configuration for testing. Use constant sampling to sample every trace
	// and enable LogSpan to log every span via configured Logger.
	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
		},
	}

	// Example logger and metrics factory. Use github.com/uber/jaeger-client-go/log
	// and github.com/uber/jaeger-lib/metrics respectively to bind to real logging and metrics
	// frameworks.
	stdLogger := jaegerlog.StdLogger
	nullMetricsFactory := metrics.NullFactory

	// Initialize tracer with a logger and a metrics factory
	closer, tracerErr := cfg.InitGlobalTracer(
		serviceName,
		jaegercfg.Logger(stdLogger),
		jaegercfg.Metrics(nullMetricsFactory),
	)
	if tracerErr != nil {
		logging.SugaredLog.Errorf("Tracer creation failed: %s", tracerErr.Error())
		os.Exit(501)
	}

	logging.SugaredLog.Infof("Global tracer registered: %t", opentracing.IsGlobalTracerRegistered())

	return closer
}

func InitProductionTracing(serviceName string) io.Closer {
	// Recommended configuration for production.
	cfg := jaegercfg.Configuration{}

	// Example logger and metrics factory. Use github.com/uber/jaeger-client-go/log
	// and github.com/uber/jaeger-lib/metrics respectively to bind to real logging and metrics
	// frameworks.
	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	// Initialize tracer with a logger and a metrics factory
	closer, tracerErr := cfg.InitGlobalTracer(
		serviceName,
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)
	if tracerErr != nil {
		logging.SugaredLog.Errorf("Tracer creation failed: %s", tracerErr.Error())
		os.Exit(501)
	}

	logging.SugaredLog.Infof("Global tracer registered: %t", opentracing.IsGlobalTracerRegistered())

	return closer
}
