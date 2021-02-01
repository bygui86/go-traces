package worker

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"

	"github.com/bygui86/go-traces/standalone/commons"
	"github.com/bygui86/go-traces/standalone/logging"
)

func (w *Worker) startWorking() {
	logging.SugaredLog.Infof("Start working every %s sec", w.config.workingInterval.String())
	w.ticker = time.NewTicker(w.config.workingInterval)

	for {
		select {
		case <-w.ticker.C:
			span := opentracing.StartSpan("simple-operation-main")
			span.SetTag("app", commons.ServiceName)
			doWork(opentracing.ContextWithSpan(context.Background(), span))
			span.Finish()

		case <-w.ctx.Done():
			return
		}
	}
}

func doWork(ctx context.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx, "simple-operation-doWork")
	defer span.Finish()

	// sleep to simulate work
	random := 0.5 + rand.Float64()*(4.5-0.5)
	time.Sleep(time.Duration(random) * time.Second)

	logging.Log.Info("Work done")

	var tracingMsg string
	traceID, sampled := extractSampledTraceID(ctx)
	if !sampled {
		tracingMsg = "traceID not sampled"
	} else {
		tracingMsg = fmt.Sprintf("traceID %s", traceID)
	}
	logging.SugaredLog.Debugf("%s", tracingMsg)
}

// ExtractSampledTraceID works like ExtractTraceID but the returned bool is only true if the returned trace id is sampled.
// copied from https://github.com/weaveworks/common/blob/master/middleware/http_tracing.go
func extractSampledTraceID(ctx context.Context) (string, bool) {
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
