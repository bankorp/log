package log

import (
	"log/slog"
	"os"
)

type Logger struct {
	internalLogger slog.Logger
}

func New() *Logger {
	internalLogger := *slog.New(
		slog.NewJSONHandler(os.Stdout, nil))

	return &Logger{
		internalLogger: internalLogger}
}

func (l *Logger) Info(message string) {
	l.internalLogger.Info(message)
}
