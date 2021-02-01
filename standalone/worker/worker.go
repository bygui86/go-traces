package worker

import (
	"context"
	"fmt"
	"time"

	"github.com/bygui86/go-traces/standalone/logging"
)

func New() (*Worker, error) {
	logging.Log.Info("Create new worker")

	cfg := loadConfig()

	return &Worker{
		config:  cfg,
		ctx:     context.Background(),
		running: false,
	}, nil
}

func (w *Worker) Start() error {
	logging.Log.Info("Start worker")

	if !w.running {
		go w.startWorking()
		w.running = true
		logging.SugaredLog.Infof("worker started")
		return nil
	}

	return fmt.Errorf("worker start failed: worker not initialized or already running")
}

func (w *Worker) Shutdown(timeout int) {
	logging.SugaredLog.Warnf("Shutdown worker, timeout %d", timeout)

	if w.ticker != nil && w.running {
		w.ticker.Stop()
		w.ctx.Done()

		time.Sleep(time.Duration(timeout) * time.Second)

		w.running = false
		return
	}

	logging.Log.Error("worker shutdown failed: worker not initialized or not running")
}
