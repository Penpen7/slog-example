package log

import (
	"log/slog"
	"os"
)

func SetDefaultLogger() {
	l := slog.New(newHandler(os.Stdout, slog.LevelInfo))
	slog.SetDefault(l)
}
