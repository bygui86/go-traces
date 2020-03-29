package monitoring

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/bygui86/go-traces/logging"
)

func newRouter() *mux.Router {
	logging.SugaredLog.Debugf("Setup new monitoring router")

	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/metrics", promhttp.Handler())
	return router
}

func newHTTPServer(host string, port int, router *mux.Router) *http.Server {
	logging.SugaredLog.Debugf("Setup new monitoring HTTP server on port %d...", port)

	return &http.Server{
		Addr:    host + ":" + strconv.Itoa(port),
		Handler: router,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}
}
