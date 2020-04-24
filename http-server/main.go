package main

import (
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bygui86/go-traces/http-server/logging"
	"github.com/bygui86/go-traces/http-server/monitoring"
	"github.com/bygui86/go-traces/http-server/rest"
	"github.com/bygui86/go-traces/http-server/tracing"
)

const (
	serviceName = "go-traces_http-server"
)

var (
	monitoringServer *monitoring.Server
	tracingCloser    io.Closer
	restServer       *rest.Server
)

func main() {
	initLogging()

	logging.Log.Info("Start go-traces")

	monitoringServer = startMonitoringServer()

	tracingCloser = initTracing()

	restServer = startRestServer()

	logging.Log.Info("go-traces up&running")

	startSysCallChannel()

	shutdownAndWait(3)
}

func initLogging() {
	err := logging.InitGlobalLogger()
	if err != nil {
		logging.SugaredLog.Errorf("Logging setup failed: %s", err.Error())
		os.Exit(501)
	}
}

func startMonitoringServer() *monitoring.Server {
	server := monitoring.NewServer()
	logging.Log.Debug("Monitoring server successfully created")

	server.Start()
	logging.Log.Debug("Monitoring successfully started")

	return server
}

func initTracing() io.Closer {
	return tracing.InitTestingTracing(serviceName)
}

func startRestServer() *rest.Server {
	logging.Log.Debug("Start REST server")

	server, err := rest.NewServer()
	if err != nil {
		logging.SugaredLog.Errorf("REST server creation failed: %s", err.Error())
		os.Exit(501)
	}
	logging.Log.Debug("REST server successfully created")

	server.Start()
	logging.Log.Debug("REST server successfully started")

	rest.RegisterCustomMetrics()

	return server
}

func startSysCallChannel() {
	syscallCh := make(chan os.Signal)
	signal.Notify(syscallCh, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	<-syscallCh
}

func shutdownAndWait(timeout int) {
	logging.SugaredLog.Warnf("Termination signal received! Timeout %d", timeout)

	if restServer != nil {
		restServer.Shutdown(timeout)
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
