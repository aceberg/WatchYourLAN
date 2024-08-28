package web

import (
	"log/slog"

	"github.com/aceberg/WatchYourLAN/internal/db"
)

func updateRoutines() {
	slog.Debug("Restarting scan routine")

	close(quitScan)

	setLogLevel()

	db.SetCurrent(appConfig)
	db.Create()

	allHosts = db.Select("now")

	quitScan = make(chan bool)
	go startScan(quitScan) // scan-routine.go
}

func setLogLevel() {
	var level slog.Level

	switch appConfig.LogLevel {
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
		slog.SetLogLoggerLevel(slog.LevelInfo)
	}
	slog.SetLogLoggerLevel(level)
}
