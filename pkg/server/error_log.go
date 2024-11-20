package server

import (
	stdlog "log"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type errorLogWriter struct {
	logger *zerolog.Logger
}

func newErrorLogLogger(logger *zerolog.Logger) *stdlog.Logger {
	return stdlog.New(&errorLogWriter{logger: logger}, "", 0)
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
