package rest

import (
	"github.com/bygui86/go-traces/http-server/logging"
	"github.com/bygui86/go-traces/http-server/utils"
)

const (
	restHostEnvVar = "REST_HOST"
	restPortEnvVar = "REST_PORT"

	restHostDefault = "0.0.0.0"
	restPortDefault = 8080
)

func loadConfig() *config {
	logging.Log.Debug("Load REST configurations")
	return &config{
		restHost: utils.GetStringEnv(restHostEnvVar, restHostDefault),
		restPort: utils.GetIntEnv(restPortEnvVar, restPortDefault),
	}
}
