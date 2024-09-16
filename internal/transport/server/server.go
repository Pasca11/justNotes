package server

import (
	"context"
	"github.com/go-chi/chi/v5"
	"net"
	"net/http"
	"time"
)

type Config struct {
	ConnType     string `yaml:"conn_type"`
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	ReadTimeout  int64  `yaml:"read_timeout"`
	WriteTimeout int64  `yaml:"write_timeout"`
	IdleTimeout  int64  `yaml:"idle_timeout"`
}

type Server struct {
	cfg *Config
	srv *http.Server
	r   chi.Router
}

func New(cfg *Config, r chi.Router) *Server {
	return &Server{
		cfg: cfg,
		r:   r,
	}
}

func (s *Server) Run() error {
	ln, err := net.Listen(s.cfg.ConnType, net.JoinHostPort("", s.cfg.Port))
	if err != nil {
		return err
	}
	s.srv = &http.Server{
		Addr:         ln.Addr().String(),
		Handler:      s.r,
		ReadTimeout:  s.setReadTimeout(),
		WriteTimeout: s.setWriteTimeout(),
		IdleTimeout:  s.setIdleTimeout(),
	}
	return s.srv.Serve(ln)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

func (s *Server) setReadTimeout() time.Duration {
	return time.Duration(s.cfg.ReadTimeout) * time.Second
}

func (s *Server) setWriteTimeout() time.Duration {
	return time.Duration(s.cfg.WriteTimeout) * time.Second
}

func (s *Server) setIdleTimeout() time.Duration {
	return time.Duration(s.cfg.IdleTimeout) * time.Second
}
