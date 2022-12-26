package db

import (
	"database/sql"
	
	// Import sqlite module
	_ "modernc.org/sqlite"

	"github.com/aceberg/WatchYourLAN/internal/check"
)

func dbExec(path, sqlStatement string) {
	db, err := sql.Open("sqlite", path)
	check.IfError(err)
	defer db.Close()

	_, err = db.Exec(sqlStatement)
	check.IfError(err)
}
