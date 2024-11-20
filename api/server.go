package api

import (
	"fmt"
	"net/http"

	connectcors "connectrpc.com/cors"
	"github.com/anuragkumar19/connect/api/gen/auth/v1/authv1connect"
	"github.com/anuragkumar19/connect/database"
	"github.com/anuragkumar19/connect/infra/nats"
	"github.com/anuragkumar19/connect/infra/smtp"
	"github.com/anuragkumar19/connect/infra/storage"
	"github.com/anuragkumar19/connect/pkg/server"
	"github.com/rs/cors"
)

type Server struct {
	*server.Server
	config *Config
}

func NewServer(config *Config, queries *database.Queries, nt *nats.NATS, s *storage.Storage, mailerClient *smtp.SMTP) (Server, error) {
	if err := config.Validate(); err != nil {
		return Server{}, fmt.Errorf("invalid metrics server config: %w", err)
	}

	// TODO: connect otel library has metrics using opentelemetry built in, translate those to prometheus
	// TODO: opentelemetry connect interceptor register

	mux := http.NewServeMux()

	{
		path, handler := authv1connect.NewRegistrationServiceHandler(authv1connect.UnimplementedRegistrationServiceHandler{})
		mux.Handle(path, handler)
	}

	srv, err := server.New(&server.Config{
		Name:           "api",
		Host:           config.Host,
		Port:           config.Port,
		HandlerTimeout: config.HandlerTimeout,
	}, withCORS(mux))
	if err != nil {
		return Server{}, fmt.Errorf("failed to create server.Server for metrics server: %w", err)
	}

	return Server{
		Server: &srv,
		config: config,
	}, nil
}

// withCORS adds CORS support to a Connect HTTP handler.
func withCORS(h http.Handler) http.Handler {
	middleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: connectcors.AllowedMethods(),
		AllowedHeaders: connectcors.AllowedHeaders(),
		ExposedHeaders: connectcors.ExposedHeaders(),
	})
	return middleware.Handler(h)
}
