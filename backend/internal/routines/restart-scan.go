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

	slog.Info("Restarting scan routine")
	setLogLevel()

	quitScan = make(chan bool)
	go startScan(quitScan) // scan-routine.go
}

func setLogLevel() {
	var level slog.Level

	slog.Info("Log level: " + conf.AppConfig.LogLevel)

	switch conf.AppConfig.LogLevel {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		slog.Error("Invalid log level. Setting default level INFO")
		level = slog.LevelInfo
	}
	slog.SetLogLoggerLevel(level)
}
