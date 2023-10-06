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

func dbExec(path, sqlStatement string) {

	mu.Lock()
	db, err := sqlx.Connect("sqlite", path)
	check.IfError(err)
	defer db.Close()

	_, err = db.Exec(sqlStatement)
	mu.Unlock()

	check.IfError(err)
}

// Select - select all hosts
func Select(path string) (dbHosts []models.Host) {

	sqlStatement := `SELECT * FROM "now" ORDER BY DATE DESC`

	mu.Lock()
	db, _ := sqlx.Connect("sqlite", path)
	defer db.Close()

	err := db.Select(&dbHosts, sqlStatement)
	mu.Unlock()

	check.IfError(err)

	return dbHosts
}

// SelectHist - select all history
func SelectHist(path string) (hist []models.History) {

	sqlStatement := `SELECT * FROM "history" ORDER BY ID DESC`

	mu.Lock()
	db, _ := sqlx.Connect("sqlite", path)
	defer db.Close()

	err := db.Select(&hist, sqlStatement)
	mu.Unlock()

	check.IfError(err)

	return hist
}
