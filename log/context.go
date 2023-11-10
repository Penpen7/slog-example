package log

import (
	"context"
)

var traceIDKey = struct{}{}

func WithTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceIDKey, traceID)
}

func getTraceID(ctx context.Context) (string, bool) {
	v, ok := ctx.Value(traceIDKey).(string)
	return v, ok
}
