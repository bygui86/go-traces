package main

import (
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bygui86/go-traces/logging"
	"github.com/bygui86/go-traces/monitoring"
	"github.com/bygui86/go-traces/rest"
	"github.com/bygui86/go-traces/tracing"
)

const (
	serviceName = "go-traces"
)

var (
	monitoringServer *monitoring.Server
	tracingCloser    io.Closer
	restServer       *rest.Server
)

/*
	TODO list
		- [x] rest server with database
		- [x] logging integration
		- [x] monitoring integration
		- [x] tracing integration
		- [x] simple traces example
		- [ ] propagated traces example
		- [x] simple traces with logging
		- [x] simple traces with monitoring
		- [ ] improve monitoring packaged
			- [ ] config (see rest)
			- [ ] logging (see rest)
			- [ ] new server (see rest)
			- [ ] start go-routine (see rest)
			- [ ] shutdown (see rest)
		- [x] readme
		- [x] makefile
		- [x] dotenv
		- [x] postman
*/
func main() {
	initLogging()

	logging.Log.Info("Start go-traces")

	monitoringServer = startMonitoringServer()

	tracingCloser = tracing.InitTestingTracing(serviceName)

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

	restServer.Shutdown(timeout)

	tracingErr := tracingCloser.Close()
	if tracingErr != nil {
		logging.SugaredLog.Errorf("Tracing closure failed: %s", tracingErr.Error())
	}

	monitoringServer.Shutdown()

	time.Sleep(time.Duration(timeout+1) * time.Second)
}
