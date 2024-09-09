package logger

import (
	"context"
	"log/slog"
	"os"
)

type logLevel string

const (
	LOCAL = "local"
	PROD  = "prod"
)

type Config struct {
	Level  logLevel
	Source bool
}

type Logger interface {
	DebugContext(ctx context.Context, msg string, args ...interface{})
	InfoContext(ctx context.Context, msg string, args ...interface{})
	WarnContext(ctx context.Context, msg string, args ...interface{})
	ErrorContext(ctx context.Context, msg string, args ...interface{})
}

func New(cfg Config) Logger {
	opts := slog.HandlerOptions{
		AddSource: cfg.Source,
	}
	switch cfg.Level {
	case LOCAL:
		opts.Level = slog.LevelInfo
	case PROD:
		opts.Level = slog.LevelWarn
	}

	log := slog.New(slog.NewTextHandler(os.Stdout, &opts))
	return log
}
