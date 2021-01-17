package rest

import (
	"github.com/bygui86/go-traces/http-client/logging"
	"github.com/bygui86/go-traces/http-client/utils"
)

const (
	restServerHostEnvVar = "REST_SERVER_HOST"
	restServerPortEnvVar = "REST_SERVER_PORT"
	restHostEnvVar       = "REST_HOST"
	restPortEnvVar       = "REST_PORT"

	restServerHostDefault = "localhost"
	restServerPortDefault = 8080
	restHostDefault       = "0.0.0.0"
	restPortDefault       = 8080
)

func loadConfig() *config {
	logging.Log.Debug("Load REST configurations")
	return &config{
		restServerHost: utils.GetStringEnv(restServerHostEnvVar, restServerHostDefault),
		restServerPort: utils.GetIntEnv(restServerPortEnvVar, restServerPortDefault),
		restHost:       utils.GetStringEnv(restHostEnvVar, restHostDefault),
		restPort:       utils.GetIntEnv(restPortEnvVar, restPortDefault),
	}
}
