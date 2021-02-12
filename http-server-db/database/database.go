package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ExpansiveWorlds/instrumentedsql"
	instrumentedsqlopentracing "github.com/ExpansiveWorlds/instrumentedsql/opentracing"
	"github.com/cenkalti/backoff"
	"github.com/lib/pq"

	"github.com/bygui86/go-traces/http-server-db/logging"
)

const (
	// no tracing
	dbConnectionStringFormat = "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s"
	dbDriverName             = "postgres"

	// with tracing
	dbConnectionString       = "postgres://%s:%s@%s:%d/%s?sslmode=%s"
	instrumentedDbDriverName = "instrumeted-" + dbDriverName

	defaultPingMaxRetry = 10
)

func New() (*sql.DB, error) {
	logging.Log.Info("Create new DB connector")

	cfg := loadConfig()

	connString := fmt.Sprintf(dbConnectionStringFormat,
		cfg.dbHost, cfg.dbPort,
		cfg.dbUsername, cfg.dbPassword, cfg.dbName,
		cfg.dbSslMode,
	)
	logging.SugaredLog.Debugf("DB connection string: %s", connString)

	return sql.Open(dbDriverName, connString)
}

func NewWithWrappedTracing() (*sql.DB, error) {
	logging.Log.Info("Create new DB connector with tracing")

	cfg := loadConfig()

	// Get a database driver.Connector for a fixed configuration.
	connString := fmt.Sprintf(dbConnectionString,
		cfg.dbUsername, cfg.dbPassword,
		cfg.dbHost, cfg.dbPort,
		cfg.dbName, cfg.dbSslMode,
	)
	logging.SugaredLog.Debugf("DB connection string: %s", connString)

	connector, connErr := pq.NewConnector(connString)
	if connErr != nil {
		return nil, connErr
	}

	sql.Register(
		instrumentedDbDriverName,
		instrumentedsql.WrapDriver(
			connector.Driver(),
			instrumentedsql.WithTracer(instrumentedsqlopentracing.NewTracer()),
			instrumentedsql.WithLogger(
				instrumentedsql.LoggerFunc(func(ctx context.Context, msg string, keyvals ...interface{}) {
					logging.SugaredLog.Infof("%s %v", msg, keyvals)
				})),
		),
	)

	return sql.OpenDB(connector), nil
}

// WARN: does not work
// imports
// 		"contrib.go.opencensus.io/integrations/ocsql"
// func NewWithOcsqlTracing() (*sql.DB, error) {
// 	logging.Log.Info("Create new DB connector with tracing")
//
// 	cfg := loadConfig()
//
// 	var connector driver.Connector
// 	var connErr error
//
// 	// Get a database driver.Connector for a fixed configuration.
// 	connector, connErr = pq.NewConnector(fmt.Sprintf(dbConnectionString,
// 		cfg.dbUsername, cfg.dbPassword,
// 		cfg.dbHost, cfg.dbPort,
// 		cfg.dbName, cfg.dbSslMode,
// 	))
// 	if connErr != nil {
// 		return nil, connErr
// 	}
//
// 	// Wrap the driver.Connector with ocsql.
// 	connector = ocsql.WrapConnector(connector, ocsql.WithAllTraceOptions())
//
// 	// Use the wrapped driver.Connector.
// 	db := sql.OpenDB(connector)
//
// 	_, tableErr := db.Exec(createTableQuery)
// 	if tableErr != nil {
// 		return nil, tableErr
// 	}
//
// 	return db, nil
// }

func InitDb(db *sql.DB) error {
	logging.Log.Info("Initialize DB")

	result, tableErr := db.Exec(createTableQuery)
	if tableErr != nil {
		return tableErr
	}
	logging.SugaredLog.Debugf("Initializion result: %s", result)
	return nil
}

func PingDb(db *sql.DB, maxRetry uint64) error {
	if maxRetry <= 0 {
		logging.SugaredLog.Warnf("PingDB maxRetry value not valid, falling back to default (%d)", defaultPingMaxRetry)
		maxRetry = defaultPingMaxRetry
	}

	// WARN: connection takes a bit time to be opened, golang application is so fast that the first ping could easily fail
	pingErr := backoff.Retry(
		func() error {
			err := db.Ping()
			if err != nil {
				logging.Log.Info("PostgreSQL connection not ready, backing off...")
				return err
			}
			logging.Log.Info("PostgreSQL connection ready")
			return nil
		},
		backoff.WithMaxRetries(backoff.NewExponentialBackOff(), maxRetry),
	)
	return pingErr
}
