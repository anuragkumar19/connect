package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

// TODO: http.Handler with tracing and metrics

var (
	ErrServerNotStarted     = errors.New("server already started")
	ErrServerAlreadyStarted = errors.New("server not started")
)

type Server struct {
	logger *zerolog.Logger
	mutex  sync.Mutex
	config *Config
	server *http.Server
	h      http.Handler
}

func New(config *Config, handler http.Handler) (Server, error) {
	if err := config.Validate(); err != nil {
		return Server{}, fmt.Errorf("invalid server config: %w", err)
	}

	logger := log.With().Str("server_name", config.Name).Logger()
	return Server{
		config: config,
		h:      handler,
		logger: &logger,
		mutex:  sync.Mutex{},
		server: nil,
	}, nil
}

func (s *Server) Start() error {
	s.mutex.Lock()
	if s.server != nil {
		return ErrServerAlreadyStarted
	}
	addr := fmt.Sprintf("%s:%s", s.config.Host, s.config.Port)
	s.server = &http.Server{
		Handler:  h2c.NewHandler(middleware.Timeout(s.config.HandlerTimeout)(recoverer(s.h)), &http2.Server{}),
		Addr:     addr,
		ErrorLog: newErrorLogLogger(s.logger),
	}
	s.mutex.Unlock()

	s.logger.Info().Str("addr", addr).Str("port", s.config.Port).Msg("starting server")
	err := s.server.ListenAndServe()
	if err != nil && errors.Is(err, http.ErrServerClosed) {
		return nil
	}
	return err
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.server == nil {
		return ErrServerNotStarted
	}
	s.logger.Info().Msg("shutting down server")
	if err := s.server.Shutdown(ctx); err != nil {
		s.logger.Error().Err(err).Msg("failed to shutdown server")
		return err
	}
	s.server = nil
	return nil
}
