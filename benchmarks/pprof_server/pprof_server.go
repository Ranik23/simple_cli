package pprof_server

import (
	"cli/config"
	"context"
	"log"
	"net/http"
	_ "net/http/pprof"
)

type PprofServer struct {
	addr string
	server *http.Server
	cfg *config.Config
}

func NewPprofServer(addr string, cfg *config.Config) *PprofServer {
	return &PprofServer{
		addr: addr,
		server: &http.Server{Addr: addr},
		cfg: cfg,
	}
}

func (s *PprofServer) Run() {
	if s.cfg.Env == "test" {
		go func() {
			log.Println("Starting Pprof Server on", s.addr)
			if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("Pprof Server failed: %v", err)
			}
		}()
	}
}

func (s *PprofServer) Stop(ctx context.Context) error {
	log.Println("Stopping Pprof Server")
	return s.server.Shutdown(ctx)
}