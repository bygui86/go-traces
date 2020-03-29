package monitoring

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Config struct {
	RestHost        string
	RestPort        int
	ShutdownTimeout int
}

type Server struct {
	Config     *Config
	Router     *mux.Router
	HTTPServer *http.Server
}
