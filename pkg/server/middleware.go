package server

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/anuragkumar19/connect/pkg/stacktrace"
	"github.com/rs/zerolog/log"
)

func recoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {
				if rvr == http.ErrAbortHandler {
					// we don't recover http.ErrAbortHandler so the response
					// to the client is aborted, this should not be logged
					panic(rvr)
				}

				log.Error().Ctx(r.Context()).RawJSON("stack", stacktrace.Marshal(debug.Stack())).Err(fmt.Errorf("handler panicked with value: %v", rvr)).Msg("panic recovered")
				if r.Header.Get("Connection") != "Upgrade" {
					w.WriteHeader(http.StatusInternalServerError)
				}
			}
		}()

		next.ServeHTTP(w, r)
	})
}
