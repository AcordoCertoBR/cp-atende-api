package logger

import (
	"log/slog"
	"os"
	"strings"

	"github.com/AcordoCertoBR/cp-atende-api/libs/config"
)

var Logger *slog.Logger

func SetupLogger(cfg *config.Config) {
	hostname, err := os.Hostname()
	if err != nil {
		slog.Error("Error obtaining host name", "error", err)
	}

	lvl := slog.LevelDebug
	switch strings.ToLower(cfg.InternalConfig.LogLevel) {
	case "debug":
		lvl = slog.LevelDebug
	case "info":
		lvl = slog.LevelInfo
	case "warn":
		lvl = slog.LevelWarn
	case "error":
		lvl = slog.LevelError
	}

	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     lvl,
		AddSource: true,
	})

	slogLogger := slog.New(handler)
	slogLogger = slogLogger.With(slog.String("app", cfg.InternalConfig.AppName), slog.String("hostname", hostname))

	slog.SetDefault(slogLogger)
}
