package main

import (
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/openzipkin/zipkin-go/reporter"

	"github.com/bygui86/go-traces/standalone/commons"
	"github.com/bygui86/go-traces/standalone/config"
	"github.com/bygui86/go-traces/standalone/logging"
	"github.com/bygui86/go-traces/standalone/monitoring"
	"github.com/bygui86/go-traces/standalone/tracing"
	"github.com/bygui86/go-traces/standalone/worker"
)

const (
	zipkinHost = "localhost"
	zipkinPort = 9411
)

var (
	monitoringServer *monitoring.Server
	jaegerCloser     io.Closer
	zipkinReporter   reporter.Reporter
	workerThread     *worker.Worker
)

func main() {
	initLogging()

	logging.SugaredLog.Infof("Start %s", commons.ServiceName)

	cfg := loadConfig()

	if cfg.GetEnableMonitoring() {
		monitoringServer = startMonitoringServer()
	}

	if cfg.GetEnableTracing() {
		switch cfg.GetTracingTech() {
		case config.TracingTechJaeger:
			jaegerCloser = initJaegerTracer()
		case config.TracingTechZipkin:
			zipkinReporter = initZipkinTracer()
		}
	}

	workerThread = startWorker()

	logging.SugaredLog.Infof("%s up and running", commons.ServiceName)

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

func loadConfig() *config.Config {
	logging.Log.Debug("Load configurations")
	return config.LoadConfig()
}

func startMonitoringServer() *monitoring.Server {
	logging.Log.Debug("Start monitoring")
	server := monitoring.New()
	logging.Log.Debug("Monitoring server successfully created")

	server.Start()
	logging.Log.Debug("Monitoring successfully started")

	return server
}

func initJaegerTracer() io.Closer {
	logging.Log.Debug("Init Jaeger tracer")
	closer, err := tracing.InitTracer()
	if err != nil {
		logging.SugaredLog.Errorf("Jaeger tracer setup failed: %s", err.Error())
		os.Exit(501)
	}
	return closer
}

func initZipkinTracer() reporter.Reporter {
	logging.Log.Debug("Init Zipkin tracer")
	zReporter, err := tracing.InitTestingZipkin(commons.ServiceName, zipkinHost, zipkinPort)
	if err != nil {
		logging.SugaredLog.Errorf("Zipkin tracer setup failed: %s", err.Error())
		os.Exit(501)
	}
	return zReporter
}

func startWorker() *worker.Worker {
	logging.Log.Debug("Start worker")

	work, newErr := worker.New()
	if newErr != nil {
		logging.SugaredLog.Errorf("Worker creation failed: %s", newErr.Error())
		os.Exit(501)
	}
	logging.Log.Debug("Worker successfully created")

	startErr := work.Start()
	if startErr != nil {
		logging.SugaredLog.Errorf("Worker start failed: %s", startErr.Error())
		os.Exit(501)
	}
	logging.Log.Debug("Worker successfully started")

	return work
}

func startSysCallChannel() {
	syscallCh := make(chan os.Signal)
	signal.Notify(syscallCh, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	<-syscallCh
}

func shutdownAndWait(timeout int) {
	logging.SugaredLog.Warnf("Termination signal received! Timeout %d", timeout)

	if jaegerCloser != nil {
		err := jaegerCloser.Close()
		if err != nil {
			logging.SugaredLog.Errorf("Jaeger tracer closure failed: %s", err.Error())
		}
	}

	if zipkinReporter != nil {
		err := zipkinReporter.Close()
		if err != nil {
			logging.SugaredLog.Errorf("Zipkin tracer closure failed: %s", err.Error())
		}
	}

	if monitoringServer != nil {
		monitoringServer.Shutdown(timeout)
	}

	time.Sleep(time.Duration(timeout+1) * time.Second)
}
