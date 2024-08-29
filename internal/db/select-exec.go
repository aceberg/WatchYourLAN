package db

import (
	// "log/slog"
	"sync"

	"github.com/aceberg/WatchYourLAN/internal/check"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

var mu sync.Mutex

func dbExec(sqlStatement string) {

	db, ok := connectDB()
	defer db.Close()

	if ok {
		mu.Lock()
		_, err := db.Exec(sqlStatement)
		mu.Unlock()
		check.IfError(err)
	}
}

// Select - get all hosts
func Select(table string) (dbHosts []models.Host) {

	sqlStatement := `SELECT * FROM ` + table + ` ORDER BY "DATE" DESC`

	db, ok := connectDB()
	defer db.Close()

	if ok {
		mu.Lock()
		err := db.Select(&dbHosts, sqlStatement)
		mu.Unlock()
		check.IfError(err)
	}

	return dbHosts
}
