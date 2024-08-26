package db

import (
	// "log/slog"

	"github.com/jmoiron/sqlx"

	// import postgres
	_ "github.com/lib/pq"

	"github.com/aceberg/WatchYourLAN/internal/check"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

func dbExecPG(sqlStatement string) {

	// slog.Debug("PG Exec " + sqlStatement)

	db, err := sqlx.Connect("postgres", currentDB.PGConnect)
	check.IfError(err)
	defer db.Close()

	_, err = db.Exec(sqlStatement)
	check.IfError(err)
}

func selectPG(table string) (dbHosts []models.Host) {

	sqlStatement := `SELECT * FROM ` + table + ` ORDER BY "DATE" DESC`

	// slog.Debug("PG Select " + sqlStatement)

	db, _ := sqlx.Connect("postgres", currentDB.PGConnect)
	defer db.Close()

	err := db.Select(&dbHosts, sqlStatement)
	check.IfError(err)

	return dbHosts
}
