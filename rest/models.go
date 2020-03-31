package rest

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

type config struct {
	dbHost     string
	dbPort     int
	dbUsername string
	dbPassword string
	dbName     string
	dbSslMode  string
	restHost   string
	restPort   int
}

type Server struct {
	config       *config
	router       *mux.Router
	httpServer   *http.Server
	dbConnection *sql.DB
	running      bool
}
