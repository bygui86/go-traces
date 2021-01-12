package rest

import (
	"github.com/bygui86/go-traces/http-server-db/logging"
	"github.com/bygui86/go-traces/http-server-db/utils"
)

const (
	restHostEnvVar = "REST_HOST"
	restPortEnvVar = "REST_PORT"

	restHostEnvVarDefault = "localhost"
	restPortEnvVarDefault = 8080
)

func loadConfig() *config {
	logging.Log.Debug("Load REST configurations")
	return &config{
		restHost: utils.GetStringEnv(restHostEnvVar, restHostEnvVarDefault),
		restPort: utils.GetIntEnv(restPortEnvVar, restPortEnvVarDefault),
	}
}
