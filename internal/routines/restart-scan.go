package routines

import (
	"log/slog"

	"github.com/aceberg/WatchYourLAN/internal/conf"
)

var (
	quitScan = make(chan bool)
)

// ScanRestart - start or update routines
func ScanRestart() {

	close(quitScan)

	slog.Debug("Restarting scan routine")
	setLogLevel()

	quitScan = make(chan bool)
	go startScan(quitScan) // scan-routine.go
}

func setLogLevel() {
	var level slog.Level

	switch conf.AppConfig.LogLevel {
	case "debug":
		slog.Info("Log level=DEBUG")
		level = slog.LevelDebug
	case "info":
		slog.Info("Log level=INFO")
		level = slog.LevelInfo
	case "warn":
		slog.Info("Log level=WARN")
		level = slog.LevelWarn
	case "error":
		slog.Info("Log level=ERROR")
		level = slog.LevelError
	default:
		slog.Error("Invalid log level. Setting default level INFO")
		level = slog.LevelInfo
	}
	slog.SetLogLoggerLevel(level)
}
