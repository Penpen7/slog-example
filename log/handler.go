package log

import (
	"context"
	"io"
	"log/slog"
)

type handler struct {
	*slog.JSONHandler
}

func newHandler(writer io.Writer, level slog.Leveler) *handler {
	return &handler{
		JSONHandler: slog.NewJSONHandler(writer, &slog.HandlerOptions{
			AddSource:   true,
			ReplaceAttr: replaceAttr,
			Level:       level,
		}),
	}
}

func (h *handler) Handle(ctx context.Context, rec slog.Record) error {
	if traceID, ok := getTraceID(ctx); ok {
		rec.AddAttrs(slog.String("trace_id", traceID))
	}
	return h.JSONHandler.Handle(ctx, rec)
}

var _ slog.Handler = (*handler)(nil)

func replaceAttr(groups []string, a slog.Attr) slog.Attr {
	switch v := a.Value.Any().(type) {
	case error:
		if stackTrace, err := getStackTraceFromError(v); err == nil {
			return slog.Group("error",
				slog.String("message", v.Error()),
				slog.Any("stacktrace", stackTrace))
		}
		return slog.Attr{
			Key:   "error",
			Value: slog.AnyValue(v),
		}
	default:
		return a
	}
}
