package metrics

import (
	"expvar"
	"fmt"
	"net/http"
	"net/http/pprof"

	"github.com/anuragkumar19/connect/pkg/server"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Server struct {
	*server.Server
	config *Config
}

func NewServer(config *Config) (Server, error) {
	if err := config.Validate(); err != nil {
		return Server{}, fmt.Errorf("invalid metrics server config: %w", err)
	}

	mux := http.NewServeMux()

	mux.Handle("/metrics", promhttp.Handler())

	// TODO: Metrics and traces for these, perquisite server.Server metrics and tracing handler
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	mux.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	mux.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	mux.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	mux.Handle("/debug/pprof/block", pprof.Handler("block"))
	mux.Handle("/debug/vars", expvar.Handler())

	s, err := server.New(&server.Config{
		Name:           "metrics",
		Host:           config.Host,
		Port:           config.Port,
		HandlerTimeout: config.HandlerTimeout,
	}, mux)
	if err != nil {
		return Server{}, fmt.Errorf("failed to create server.Server for metrics server: %w", err)
	}

	return Server{
		Server: &s,
		config: config,
	}, nil
}
