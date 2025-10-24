package core

import (
	"fmt"
	"net/http"
	"time"

	chimiddleware "github.com/go-chi/chi/v5/middleware"
)

// NewRequestLogger returns a chi RequestLogger middleware that emits structured
// logs using the provided application logger.
func NewRequestLogger(logger Logger) func(http.Handler) http.Handler {
	if logger == nil {
		logger = NewNoopLogger()
	}

	return chimiddleware.RequestLogger(&structuredLogFormatter{logger: logger})
}

type structuredLogFormatter struct {
	logger Logger
}

func (f *structuredLogFormatter) NewLogEntry(r *http.Request) chimiddleware.LogEntry {
	reqID := RequestIDFrom(r.Context())
	entryLogger := f.logger.With(
		"request_id", reqID,
		"method", r.Method,
		"path", r.URL.Path,
	)

	entry := &structuredLogEntry{
		logger: entryLogger,
		req:    r,
		start:  time.Now(),
	}

	entryLogger.Debug("request started",
		"remote_addr", r.RemoteAddr,
		"user_agent", r.UserAgent(),
	)

	return entry
}

type structuredLogEntry struct {
	logger Logger
	req    *http.Request
	start  time.Time
}

func (e *structuredLogEntry) Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{}) {
	e.logger.Info("request completed",
		"status", status,
		"bytes", bytes,
		"elapsed_ms", elapsed.Milliseconds(),
		"referer", e.req.Referer(),
	)
}

func (e *structuredLogEntry) Panic(v interface{}, stack []byte) {
	e.logger.Error("request panic",
		"panic", fmt.Sprint(v),
		"stack", string(stack),
	)
}
