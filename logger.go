package bot

import (
	"context"
	"os"

	"log/slog"
)

// https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry#LogSeverity
const SeverityKey string = "severity"

const (
	LevelDebug    string = "DEBUG"
	LevelInfo            = "INFO"
	LevelWarn            = "WARNING"
	LevelError           = "ERROR"
	LevelCritical        = "CRITICAL"
)

func ConvertLevelToSlogLevel(level string) slog.Level {
	switch level {
	case LevelDebug:
		return slog.LevelDebug
	case LevelInfo:
		return slog.LevelInfo
	case LevelWarn:
		return slog.LevelWarn
	case LevelError:
		return slog.LevelError
	case LevelCritical:
		return slog.LevelError
	default:
		return slog.LevelDebug
	}
}

func InitLogger(c Config) {
	opt := &slog.HandlerOptions{
		Level: ConvertLevelToSlogLevel(c.LogLevel),
		ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
			switch a.Key {
			case slog.MessageKey:
				return slog.String("message", a.Value.String())
			case slog.LevelKey:
				return slog.Attr{}
			}
			return a
		},
	}
	h := WithContextValueHandler(slog.NewJSONHandler(os.Stdout, opt))
	logger := slog.New(h)
	slog.SetDefault(logger)
}

type ContextValueHandler struct {
	parent slog.Handler
}

func WithContextValueHandler(parent slog.Handler) *ContextValueHandler {
	return &ContextValueHandler{
		parent: parent,
	}
}

func (h *ContextValueHandler) Handle(ctx context.Context, record slog.Record) error {
	record.AddAttrs(slog.String("requestId", GetRequestID(ctx)))
	return h.parent.Handle(ctx, record)
}

func (h *ContextValueHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.parent.Enabled(ctx, level)
}

func (h *ContextValueHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &ContextValueHandler{h.parent.WithAttrs(attrs)}
}

func (h *ContextValueHandler) WithGroup(name string) slog.Handler {
	return &ContextValueHandler{h.parent.WithGroup(name)}
}

func Debug(ctx context.Context, msg string) {
	slog.DebugContext(ctx, msg,
		slog.String(SeverityKey, LevelDebug),
	)
}

func Info(ctx context.Context, msg string) {
	slog.InfoContext(ctx, msg,
		slog.String(SeverityKey, LevelInfo),
	)
}

func Warn(ctx context.Context, msg string) {
	slog.WarnContext(ctx, msg,
		slog.String(SeverityKey, LevelWarn),
	)
}

func Error(ctx context.Context, msg string) {
	slog.ErrorContext(ctx, msg,
		slog.String(SeverityKey, LevelError),
	)
}

func Fatal(ctx context.Context, msg string) {
	slog.ErrorContext(ctx, msg,
		slog.String(SeverityKey, LevelCritical),
	)
}
