package rest

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/bygui86/go-traces/http-server/database"
)

type Server struct {
	config     *config
	router     *mux.Router
	httpServer *http.Server
	db         database.InMemoryDbInt
	running    bool
}

type config struct {
	restHost string
	restPort int
}
