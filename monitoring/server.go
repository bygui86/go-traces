package monitoring

import (
	"context"
	"time"

	"github.com/bygui86/go-traces/logging"
)

func NewServer() *Server {
	logging.Log.Info("Create new monitoring server")

	cfg := loadConfig()
	router := newRouter()
	httpServer := newHTTPServer(cfg.RestHost, cfg.RestPort, router)
	return &Server{
		Config:     cfg,
		Router:     router,
		HTTPServer: httpServer,
	}
}

func (s *Server) Start() {
	logging.Log.Info("Start monitoring server")

	go func() {
		if err := s.HTTPServer.ListenAndServe(); err != nil {
			logging.SugaredLog.Errorf("Monitoring server start failed: %s", err.Error())
		}
	}()

	logging.SugaredLog.Debugf("Monitoring server listen on port", s.Config.RestPort)
}

func (s *Server) Shutdown() {
	logging.Log.Warn("Shutdown monitoring server")

	if s.HTTPServer != nil {
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.Config.ShutdownTimeout)*time.Second)
		defer cancel()
		// does not block if no connections, otherwise wait until the timeout deadline
		err := s.HTTPServer.Shutdown(ctx)
		if err != nil {
			logging.SugaredLog.Errorf("Monitoring server shutdown failed: %s", err.Error())
		}
	}
}
