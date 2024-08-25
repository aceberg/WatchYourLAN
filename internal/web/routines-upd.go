package web

import (
	"log/slog"

	"github.com/aceberg/WatchYourLAN/internal/db"
)

func updateRoutines() {

	close(quitScan)

	setLogLevel()

	db.Create(appConfig.DBPath)

	allHosts = db.Select(appConfig.DBPath, "now")
	histHosts = db.Select(appConfig.DBPath, "history")

	quitScan = make(chan bool)
	go startScan(quitScan) // scan-routine.go
}

func setLogLevel() {
	var level slog.Level

	switch appConfig.LogLevel {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		// invalid log level, handle error
	}
	slog.SetLogLoggerLevel(level)
}
