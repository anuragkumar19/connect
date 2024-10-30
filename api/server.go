package api

import (
	"net/http"

	connectcors "connectrpc.com/cors"
	"github.com/anuragkumar19/connect/api/gen/auth/v1/authv1connect"
	"github.com/anuragkumar19/connect/api/services/auth/login"
	"github.com/anuragkumar19/connect/api/services/auth/passwordreset"
	"github.com/anuragkumar19/connect/api/services/auth/registration"
	"github.com/anuragkumar19/connect/database"
	"github.com/anuragkumar19/connect/infra/nats"
	"github.com/anuragkumar19/connect/infra/smtp"
	"github.com/anuragkumar19/connect/infra/storage"
	"github.com/rs/cors"
	"github.com/rs/zerolog"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type HTTPServer struct {
	registrationService  *registration.Service
	loginService         *login.Service
	passwordResetService *passwordreset.Service
}

func NewServer(logger *zerolog.Logger, db *database.Queries, nt *nats.NATS, s *storage.Storage, mailerClient *smtp.SMTP) HTTPServer {
	registrationService := registration.New(logger)
	loginService := login.New(logger)
	passwordService := passwordreset.New(logger)

	return HTTPServer{
		registrationService:  &registrationService,
		loginService:         &loginService,
		passwordResetService: &passwordService,
	}
}

func (s *HTTPServer) Serve() error {
	api := http.NewServeMux()
	{
		path, handler := authv1connect.NewRegistrationServiceHandler(s.registrationService)
		api.Handle(path, handler)
	}
	{
		path, handler := authv1connect.NewLoginServiceHandler(s.loginService)
		api.Handle(path, handler)
	}
	{
		path, handler := authv1connect.NewPasswordResetServiceHandler(s.passwordResetService)
		api.Handle(path, handler)
	}

	mux := http.NewServeMux()

	mux.Handle("/api/", http.StripPrefix("/api", api))

	return http.ListenAndServe("localhost:8080", h2c.NewHandler(withCORS(mux), &http2.Server{}))
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
