package monitoring

import (
	"github.com/bygui86/go-traces/logging"
	"github.com/bygui86/go-traces/utils"
)

const (
	monHostEnvVar            = "MONITOR_HOST"
	monPortEnvVar            = "MONITOR_PORT"
	monShutdownTimeoutEnvVar = "MONITOR_SHUTDOWN_TIMEOUT"

	monHostDefault     = "localhost"
	monPortDefault     = 9090
	monShutdownTimeout = 3
)

func loadConfig() *Config {
	logging.SugaredLog.Debug("Load monitoring configurations")
	return &Config{
		RestHost:        utils.GetStringEnv(monHostEnvVar, monHostDefault),
		RestPort:        utils.GetIntEnv(monPortEnvVar, monPortDefault),
		ShutdownTimeout: utils.GetIntEnv(monShutdownTimeoutEnvVar, monShutdownTimeout),
	}
}
