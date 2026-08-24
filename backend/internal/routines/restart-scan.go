package routines

import (
	"log/slog"
	"sync"

	"github.com/aceberg/WatchYourLAN/internal/conf"
)

var (
	quitScan      = make(chan bool)
	scanRestartMu sync.Mutex
	startScanFunc = startScan
)

// ScanRestart - start or update routines
func ScanRestart() {
	scanRestartMu.Lock()
	defer scanRestartMu.Unlock()

	close(quitScan)

	slog.Info("Restarting scan routine")
	setLogLevel()

	quitScan = make(chan bool)
	go startScanFunc(quitScan) // scan-routine.go
}

func setLogLevel() {
	var level slog.Level
	config := conf.GetAppConfig()

	slog.Info("Log level: " + config.LogLevel)

	switch config.LogLevel {
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
