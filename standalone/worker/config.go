package worker

import (
	"time"

	"github.com/bygui86/go-traces/standalone/logging"
	"github.com/bygui86/go-traces/standalone/utils"
)

const (
	workingIntervalEnvVar = "WORKING_INTERVAL" // in seconds

	workingIntervalDefault = 5
)

func loadConfig() *config {
	logging.Log.Info("Load configurations")

	return &config{
		workingInterval: time.Duration(utils.GetIntEnv(workingIntervalEnvVar, workingIntervalDefault)) * time.Second,
	}
}
