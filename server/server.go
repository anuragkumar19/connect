package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

var ErrServerNotStarted = errors.New("server already started")
var ErrServerAlreadyStarted = errors.New("server already started")

type Server struct {
	mutex  sync.Mutex
	config *Config
	logger *zerolog.Logger
	server *http.Server
	router *chi.Mux
}

func New(config *Config, logger *zerolog.Logger) *Server {
	r := chi.NewRouter()

	r.Use(recoverer(logger))
	r.Use(middleware.Timeout(config.HandlerTimeout))

	return &Server{
		config: config,
		logger: logger,
		router: r,
	}
}

func (s *Server) Start() error {
	s.mutex.Lock()
	if s.server != nil {
		return ErrServerAlreadyStarted
	}
	addr := fmt.Sprintf("%s:%s", s.config.Host, s.config.Port)
	s.server = &http.Server{
		Handler:  h2c.NewHandler(s.router, &http2.Server{}),
		Addr:     addr,
		ErrorLog: log.New(newErrorLogWriter(s.logger), "", 0),
	}
	s.mutex.Unlock()

	s.logger.Info().Str("addr", addr).Msg("starting and listening server")
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
	err := s.server.Shutdown(ctx)
	s.server = nil
	return err
}

// Mount strip the prefix and mount handler to *chi.Mux. Handles chi's panic gracefully
func (s *Server) Mount(prefix string, h http.Handler) (err error) {
	// Chi.Router.Mount panics if there is duplicate mount path
	defer func() {
		if v := recover(); v != nil {
			err = fmt.Errorf("%v", v)
		}
	}()

	s.mutex.Lock()
	if s.server != nil {
		return ErrServerAlreadyStarted
	}
	s.mutex.Unlock()
	s.router.Mount(prefix, http.StripPrefix(prefix, h))
	return err
}
