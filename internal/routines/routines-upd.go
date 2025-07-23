package routines

import (
	"log/slog"

	"github.com/aceberg/WatchYourLAN/internal/conf"
	"github.com/aceberg/WatchYourLAN/internal/db"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

var (
	allHosts  []models.Host
	histHosts []models.Host

	quitScan = make(chan bool)
)

// Update - start or update routines
func Update() {

	slog.Debug("Restarting scan routine")

	close(quitScan)

	setLogLevel()

	db.SetCurrent(conf.AppConfig)
	db.Create()

	allHosts = db.Select("now")

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
