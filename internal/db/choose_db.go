package db

import (
	"log/slog"

	"github.com/aceberg/WatchYourLAN/internal/models"
)

// Data to connect to DB
type Data struct {
	Use        string
	Path       string
	SQLitePath string
	PGConnect  string
	PrimaryKey string
}

var currentDB Data

func setCurrentDB() {

	if currentDB.Use == "postgres" && currentDB.PGConnect != "" {
		currentDB.Path = currentDB.PGConnect
		currentDB.PrimaryKey = "BIGSERIAL PRIMARY KEY"
	} else {
		currentDB.Use = "sqlite"
		currentDB.Path = currentDB.SQLitePath
		currentDB.PrimaryKey = "INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE"
	}

	slog.Info("Using DB", "type", currentDB.Use)
}

// SetCurrent - set paths and which DB to use
func SetCurrent(config models.Conf) {

	currentDB.Use = config.UseDB
	currentDB.SQLitePath = config.DBPath
	currentDB.PGConnect = config.PGConnect

	setCurrentDB()
}
