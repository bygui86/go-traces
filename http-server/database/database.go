package database

import (
	"database/sql"
	"fmt"

	"contrib.go.opencensus.io/integrations/ocsql"
	"github.com/lib/pq"
	_ "github.com/lib/pq"

	"github.com/bygui86/go-traces/http-server/logging"
)

const (
	// no tracing
	dbConnectionStringFormat = "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s"
	dbDriverName             = "postgres"

	// with tracing
	dbConnectionString = "postgres://%s:%s@%s:%d/%s?sslmode=%s"
)

func New() (*sql.DB, error) {
	logging.Log.Info("Create new DB connector")

	cfg := loadConfig()

	db, dbErr := sql.Open(
		dbDriverName,
		fmt.Sprintf(dbConnectionStringFormat,
			cfg.dbHost, cfg.dbPort,
			cfg.dbUsername, cfg.dbPassword, cfg.dbName,
			cfg.dbSslMode,
		),
	)
	if dbErr != nil {
		return nil, dbErr
	}

	_, tableErr := db.Exec(createTableQuery)
	if tableErr != nil {
		return nil, tableErr
	}

	return db, nil
}

// TODO does not work!
func NewWithTracing() (*sql.DB, error) {
	logging.Log.Info("Create new DB connector")

	cfg := loadConfig()

	// Get a database driver.Connector for a fixed configuration.
	connector, connErr := pq.NewConnector(fmt.Sprintf(dbConnectionString,
		cfg.dbUsername, cfg.dbPassword,
		cfg.dbHost, cfg.dbPort,
		cfg.dbName, cfg.dbSslMode,
	))
	if connErr != nil {
		return nil, connErr
	}

	// Wrap the driver.Connector with ocsql.
	ocConnector := ocsql.WrapConnector(connector, ocsql.WithAllTraceOptions())

	// Use the wrapped driver.Connector.
	db := sql.OpenDB(ocConnector)

	_, tableErr := db.Exec(createTableQuery)
	if tableErr != nil {
		return nil, tableErr
	}

	return db, nil
}
