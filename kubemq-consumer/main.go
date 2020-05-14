package main

import (
	"context"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bygui86/go-traces/kubemq-consumer/consumer"
	"github.com/bygui86/go-traces/kubemq-consumer/logging"
	"github.com/bygui86/go-traces/kubemq-consumer/monitoring"
	"github.com/bygui86/go-traces/kubemq-consumer/tracing"
)

const (
	serviceName = "kubemq-consumer"
)

var (
	monitoringServer *monitoring.Server
	tracingCloser    io.Closer
	ctxCancel        context.CancelFunc
	kubemqConsumer   *consumer.KubemqConsumer
)

func main() {
	initLogging()

	logging.SugaredLog.Infof("Start %s", serviceName)

	monitoringServer = startMonitoringServer()

	tracingCloser = initTracing()

	var ctx context.Context
	ctx, ctxCancel = context.WithCancel(context.Background())

	kubemqConsumer = startConsumer(ctx)

	logging.SugaredLog.Infof("%s up and running", serviceName)

	startSysCallChannel()

	shutdownAndWait(1)
}

func initLogging() {
	err := logging.InitGlobalLogger()
	if err != nil {
		logging.SugaredLog.Errorf("Logging setup failed: %s", err.Error())
		os.Exit(501)
	}
}

func startMonitoringServer() *monitoring.Server {
	logging.Log.Debug("Start monitoring")
	server := monitoring.New()
	logging.Log.Debug("Monitoring server successfully created")

	server.Start()
	logging.Log.Debug("Monitoring successfully started")

	return server
}

func initTracing() io.Closer {
	logging.Log.Debug("Init tracing")
	return tracing.InitTestingTracing(serviceName)
}

func startConsumer(ctx context.Context) *consumer.KubemqConsumer {
	logging.Log.Debug("Start consumer")
	mqConsumer, newErr := consumer.New(ctx, serviceName)
	if newErr != nil {
		logging.SugaredLog.Errorf("Consumer setup failed: %s", newErr.Error())
		os.Exit(501)
	}
	logging.Log.Debug("Consumer successfully created")

	startErr := mqConsumer.Start()
	if startErr != nil {
		logging.SugaredLog.Errorf("Consumer start failed: %s", startErr.Error())
		os.Exit(502)
	}
	logging.Log.Debug("Consumer successfully started")

	return mqConsumer
}

func startSysCallChannel() {
	syscallCh := make(chan os.Signal)
	signal.Notify(syscallCh, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	<-syscallCh
}

func shutdownAndWait(timeout int) {
	logging.SugaredLog.Warnf("Termination signal received! Timeout %d", timeout)

	if kubemqConsumer != nil {
		kubemqConsumer.Shutdown(timeout)
	}

	if ctxCancel != nil {
		ctxCancel()
	}

	if tracingCloser != nil {
		tracingErr := tracingCloser.Close()
		if tracingErr != nil {
			logging.SugaredLog.Errorf("Tracing closure failed: %s", tracingErr.Error())
		}
	}

	if monitoringServer != nil {
		monitoringServer.Shutdown(timeout)
	}

	time.Sleep(time.Duration(timeout+1) * time.Second)
}
