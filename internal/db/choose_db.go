package db

import (
	"log/slog"

	"github.com/aceberg/WatchYourLAN/internal/models"
)

// Data to connect to DB
type Data struct {
	Use        string
	SQLitePath string
	PGConnect  string
	PrimaryKey string
}

var currentDB Data

func dbExec(sqlStatement string) {

	if currentDB.Use == "postgres" {
		dbExecPG(sqlStatement)
	} else {
		dbExecLite(sqlStatement)
	}
}

// Select - select all from table
func Select(table string) (dbHosts []models.Host) {

	if currentDB.Use == "postgres" {
		dbHosts = selectPG(table)
	} else {
		dbHosts = selectLite(table)
	}

	return dbHosts
}

// SetCurrent - set paths and which DB to use
func SetCurrent(config models.Conf) {

	currentDB.Use = config.UseDB
	currentDB.SQLitePath = config.DBPath
	currentDB.PGConnect = config.PGConnect

	if currentDB.Use == "postgres" && currentDB.PGConnect != "" {
		currentDB.PrimaryKey = "BIGSERIAL PRIMARY KEY"
	} else {
		currentDB.Use = "sqlite"
		currentDB.PrimaryKey = "INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE"
	}

	slog.Info("Using DB", "type", currentDB.Use)
}
