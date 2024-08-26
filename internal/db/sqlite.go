package db

import (
	"sync"

	"github.com/jmoiron/sqlx"

	// Import sqlite module
	_ "modernc.org/sqlite"

	"github.com/aceberg/WatchYourLAN/internal/check"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

var mu sync.Mutex

func dbExecLite(sqlStatement string) {

	mu.Lock()
	db, err := sqlx.Connect("sqlite", currentDB.SQLitePath)
	check.IfError(err)
	defer db.Close()

	_, err = db.Exec(sqlStatement)
	mu.Unlock()

	check.IfError(err)
}

func selectLite(table string) (dbHosts []models.Host) {

	sqlStatement := "SELECT * FROM " + table + " ORDER BY DATE DESC"

	mu.Lock()
	db, _ := sqlx.Connect("sqlite", currentDB.SQLitePath)
	defer db.Close()

	err := db.Select(&dbHosts, sqlStatement)
	mu.Unlock()

	check.IfError(err)

	return dbHosts
}
