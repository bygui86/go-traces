package main

import (
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bygui86/go-traces/logging"
	"github.com/bygui86/go-traces/rest"
	"github.com/bygui86/go-traces/tracing"
)

const (
	serviceName = "go-traces"
)

var (
	tracingCloser io.Closer
	restServer    *rest.Server
)

func main() {
	logging.Log.Info("Start go-traces")

	setupLogging()

	// TODO start monitoring

	tracingCloser = tracing.InitSample(serviceName)

	restServer = startRestServer()

	logging.Log.Info("go-traces up&running")

	startSysCallChannel()

	shutdownAndWait(3)
}

func setupLogging() {
	err := logging.InitGlobalLogger()
	if err != nil {
		logging.SugaredLog.Errorf("Logging setup failed: %s", err.Error())
		os.Exit(501)
	}
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

	time.Sleep(time.Duration(timeout+1) * time.Second)
}
