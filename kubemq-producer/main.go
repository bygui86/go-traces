package main

import (
	"context"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bygui86/go-traces/kubemq-producer/logging"
	"github.com/bygui86/go-traces/kubemq-producer/monitoring"
	"github.com/bygui86/go-traces/kubemq-producer/producer"
	"github.com/bygui86/go-traces/kubemq-producer/tracing"
)

const (
	serviceName = "kubemq-producer"
)

var (
	monitoringServer *monitoring.Server
	tracingCloser    io.Closer
	ctxCancel        context.CancelFunc
	kubemqProducer   *producer.KubemqProducer
)

func main() {
	initLogging()

	logging.SugaredLog.Infof("Start %s", serviceName)

	monitoringServer = startMonitoringServer()

	tracingCloser = initTracing()

	var ctx context.Context
	ctx, ctxCancel = context.WithCancel(context.Background())

	kubemqProducer = startProducer(ctx)

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

func startProducer(ctx context.Context) *producer.KubemqProducer {
	logging.Log.Debug("Start producer")
	mqProducer, newErr := producer.New(ctx, serviceName)
	if newErr != nil {
		logging.SugaredLog.Errorf("Producer setup failed: %s", newErr.Error())
		os.Exit(501)
	}
	logging.Log.Debug("Producer successfully created")

	mqProducer.Start()
	logging.Log.Debug("Producer successfully started")

	return mqProducer
}

func startSysCallChannel() {
	syscallCh := make(chan os.Signal)
	signal.Notify(syscallCh, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	<-syscallCh
}

func shutdownAndWait(timeout int) {
	logging.SugaredLog.Warnf("Termination signal received! Timeout %d", timeout)

	if kubemqProducer != nil {
		kubemqProducer.Shutdown(timeout)
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
