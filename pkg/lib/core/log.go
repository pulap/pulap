package core

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	chimiddleware "github.com/go-chi/chi/v5/middleware"
)

type LogLevel int

const (
	DebugLevel LogLevel = iota
	InfoLevel
	ErrorLevel
)

type Logger interface {
	Debug(v ...any)
	Debugf(format string, a ...any)
	Info(v ...any)
	Infof(format string, a ...any)
	Error(v ...any)
	Errorf(format string, a ...any)
	SetLogLevel(level LogLevel)
	With(args ...any) Logger
}

type slogLogger struct {
	logger   *slog.Logger
	logLevel LogLevel
}

func NewLogger(logLevelStr string) Logger {
	level := toValidLevel(logLevelStr)

	var handler slog.Handler
	if isTerminal() {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slogLevel(level),
		})
	} else {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slogLevel(level),
		})
	}

	return &slogLogger{
		logger:   slog.New(handler),
		logLevel: level,
	}
}

func (l *slogLogger) Debug(v ...any) {
	if l.logLevel <= DebugLevel {
		msg, attrs := normalizeArgs(v...)
		l.logger.Debug(msg, attrs...)
	}
}

func (l *slogLogger) Debugf(format string, a ...any) {
	if l.logLevel <= DebugLevel {
		l.logger.Debug(fmt.Sprintf(format, a...))
	}
}

func (l *slogLogger) Info(v ...any) {
	if l.logLevel <= InfoLevel {
		msg, attrs := normalizeArgs(v...)
		l.logger.Info(msg, attrs...)
	}
}

func (l *slogLogger) Infof(format string, a ...any) {
	if l.logLevel <= InfoLevel {
		l.logger.Info(fmt.Sprintf(format, a...))
	}
}

func (l *slogLogger) Error(v ...any) {
	if l.logLevel <= ErrorLevel {
		msg, attrs := normalizeArgs(v...)
		l.logger.Error(msg, attrs...)
	}
}

func (l *slogLogger) Errorf(format string, a ...any) {
	if l.logLevel <= ErrorLevel {
		l.logger.Error(fmt.Sprintf(format, a...))
	}
}

func (l *slogLogger) SetLogLevel(level LogLevel) {
	l.logLevel = level
}

func (l *slogLogger) With(args ...any) Logger {
	return &slogLogger{
		logger:   l.logger.With(args...),
		logLevel: l.logLevel,
	}
}

type noopLogger struct{}

func (noopLogger) Debug(v ...any)                 {}
func (noopLogger) Debugf(format string, a ...any) {}
func (noopLogger) Info(v ...any)                  {}
func (noopLogger) Infof(format string, a ...any)  {}
func (noopLogger) Error(v ...any)                 {}
func (noopLogger) Errorf(format string, a ...any) {}
func (noopLogger) SetLogLevel(level LogLevel)     {}
func (noopLogger) With(args ...any) Logger        { return noopLogger{} }

func NewNoopLogger() Logger {
	return noopLogger{}
}

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

func toValidLevel(level string) LogLevel {
	level = strings.ToLower(level)
	switch level {
	case "debug", "dbg":
		return DebugLevel
	case "info", "inf":
		return InfoLevel
	case "error", "err":
		return ErrorLevel
	default:
		return InfoLevel
	}
}

func slogLevel(level LogLevel) slog.Level {
	switch level {
	case DebugLevel:
		return slog.LevelDebug
	case InfoLevel:
		return slog.LevelInfo
	case ErrorLevel:
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func isTerminal() bool {
	return os.Getenv("LOG_FORMAT") != "json"
}

func normalizeArgs(args ...any) (string, []any) {
	if len(args) == 0 {
		return "", nil
	}

	msg := fmt.Sprint(args[0])
	if len(args) == 1 {
		return msg, nil
	}

	rest := args[1:]
	if len(rest) == 0 {
		return msg, nil
	}

	if attrs := toAttrsIfPossible(rest); attrs != nil {
		return msg, attrs
	}

	if len(rest)%2 != 0 {
		// Fallback to plain message if key/value pairs are malformed
		return fmt.Sprint(args...), nil
	}

	return msg, rest
}

func toAttrsIfPossible(args []any) []any {
	attrs := make([]any, len(args))
	for i, arg := range args {
		if attr, ok := arg.(slog.Attr); ok {
			attrs[i] = attr
			continue
		}
		// If we encounter something that isn't an Attr, bail out
		return nil
	}
	return attrs
}
