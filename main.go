package main

import (
	"context"
	"log/slog"

	"example.com/m/log"

	"github.com/cockroachdb/errors"
)

func main() {
	log.SetDefaultLogger()
	ctx := context.Background()
	ctx = log.WithTraceID(ctx, "123")
	slog.InfoContext(ctx, "start main")
	err := generateErrorWithStackTrace()
	if err != nil {
		slog.ErrorContext(ctx, "error", err)
	}
}

func generateErrorWithStackTrace() error {
	return errors.New("this is a sample error")
}
