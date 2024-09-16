package logger

import (
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
	Info(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Debug(msg string, args ...interface{})
}

func New(cfg *Config) Logger {
	opts := slog.HandlerOptions{
		AddSource: cfg.Source,
	}
	switch cfg.Level {
	case LOCAL:
		opts.Level = slog.LevelInfo
	case PROD:
		opts.Level = slog.LevelError
	}
	opts.Level = slog.LevelInfo
	log := slog.New(slog.NewJSONHandler(os.Stdout, &opts))
	return log
}
