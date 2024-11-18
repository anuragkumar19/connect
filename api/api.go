package api

import (
	"net/http"

	connectcors "connectrpc.com/cors"
	"github.com/anuragkumar19/connect/database"
	"github.com/anuragkumar19/connect/infra/nats"
	"github.com/anuragkumar19/connect/infra/smtp"
	"github.com/anuragkumar19/connect/infra/storage"
	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
)

type Api struct {
}

func New(queries *database.Queries, nt *nats.NATS, s *storage.Storage, mailerClient *smtp.SMTP) Api {
	return Api{}
}

func (api *Api) Router() chi.Router {
	r := chi.NewRouter()
	r.Use(withCORS)

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

// func chiPath(s string) string {
// 	return s + "*"
// }
