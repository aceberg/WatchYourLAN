package db

import (
	"log/slog"

	"github.com/jmoiron/sqlx"

	// import postgres
	_ "github.com/lib/pq"

	// Import sqlite module
	_ "modernc.org/sqlite"

	"github.com/aceberg/WatchYourLAN/internal/check"
)

func connectDB() (*sqlx.DB, bool) {
	var ok bool

	db, err := sqlx.Open(currentDB.Use, currentDB.Path)
	check.IfError(err)

	err = db.Ping()
	if check.IfError(err) && currentDB.Use == "postgres" {
		slog.Warn("PostgreSQL connection error. Falling back to SQLite.")
		currentDB.Use = "sqlite"
		setCurrentDB()
		Create()
	}

	if err == nil {
		ok = true
	}

	return db, ok
}
