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
	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
	"github.com/rs/zerolog"
)

type Api struct {
	registrationService  authv1connect.RegistrationServiceHandler
	loginService         authv1connect.LoginServiceHandler
	passwordResetService authv1connect.PasswordResetServiceHandler
}

func New(logger *zerolog.Logger, store *database.Queries, nt *nats.NATS, s *storage.Storage, mailerClient *smtp.SMTP) Api {
	registrationService := registration.New(logger, store)
	loginService := login.New(logger)
	passwordResetService := passwordreset.New(logger)

	return Api{
		registrationService:  &registrationService,
		loginService:         &loginService,
		passwordResetService: &passwordResetService,
	}
}

func (api *Api) Router() chi.Router {
	r := chi.NewRouter()
	r.Use(withCORS)

	{
		path, handler := authv1connect.NewRegistrationServiceHandler(api.registrationService)
		r.Handle(chiPath(path), handler)
	}
	{
		path, handler := authv1connect.NewLoginServiceHandler(api.loginService)
		r.Handle(chiPath(path), handler)
	}
	{
		path, handler := authv1connect.NewPasswordResetServiceHandler(api.passwordResetService)
		r.Handle(chiPath(path), handler)
	}

	return r
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

func chiPath(s string) string {
	return s + "*"
}
