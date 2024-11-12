package server

import (
	"github.com/rs/zerolog"
)

type errorLogWriter struct {
	logger *zerolog.Logger
}

func newErrorLogWriter(logger *zerolog.Logger) *errorLogWriter {
	return &errorLogWriter{
		logger: logger,
	}
}

func (w *errorLogWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	if n > 0 && p[n-1] == '\n' {
		// Trim CR added by stdlog.
		p = p[0 : n-1]
	}
	w.logger.Error().Err(err).Msg(string(p))
	return
}
