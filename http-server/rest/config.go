package rest

import (
	"github.com/bygui86/go-traces/http-server/logging"
	"github.com/bygui86/go-traces/http-server/utils"
)

const (
	dbHostEnvVar     = "DB_HOST"
	dbPortEnvVar     = "DB_PORT"
	dbUsernameEnvVar = "DB_USERNAME"
	dbPasswordEnvVar = "DB_PASSWORD"
	dbNameEnvVar     = "DB_NAME"
	dbSslModeEnvVar  = "DB_SSL_MODE"
	restHostEnvVar   = "REST_HOST"
	restPortEnvVar   = "REST_PORT"

	dbHostEnvVarDefault     = "localhost"
	dbPortEnvVarDefault     = 5432
	dbUsernameEnvVarDefault = "username"
	dbPasswordEnvVarDefault = "password"
	dbNameEnvVarDefault     = "db"
	dbSslModeEnvVarDefault  = "disable"
	restHostEnvVarDefault   = "localhost"
	restPortEnvVarDefault   = 8080
)

func loadConfig() *config {
	logging.Log.Debug("Load REST configurations")
	return &config{
		dbHost:     utils.GetStringEnv(dbHostEnvVar, dbHostEnvVarDefault),
		dbPort:     utils.GetIntEnv(dbPortEnvVar, dbPortEnvVarDefault),
		dbUsername: utils.GetStringEnv(dbUsernameEnvVar, dbUsernameEnvVarDefault),
		dbPassword: utils.GetStringEnv(dbPasswordEnvVar, dbPasswordEnvVarDefault),
		dbName:     utils.GetStringEnv(dbNameEnvVar, dbNameEnvVarDefault),
		dbSslMode:  utils.GetStringEnv(dbSslModeEnvVar, dbSslModeEnvVarDefault),
		restHost:   utils.GetStringEnv(restHostEnvVar, restHostEnvVarDefault),
		restPort:   utils.GetIntEnv(restPortEnvVar, restPortEnvVarDefault),
	}
}