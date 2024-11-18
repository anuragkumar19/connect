package server

import (
	stdlog "log"

	"github.com/rs/zerolog/log"
)

var httpErrorLogger = stdlog.New(&errorLogWriter{}, "", 0)

type errorLogWriter struct {
}

func (w *errorLogWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	if n > 0 && p[n-1] == '\n' {
		// Trim CR added by stdlog.
		p = p[0 : n-1]
	}
	log.Error().Err(err).Msg(string(p))
	return
}
