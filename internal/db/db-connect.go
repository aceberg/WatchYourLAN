package db

import (
	"database/sql"
	"log"
	"sync"

	// Import sqlite module
	_ "modernc.org/sqlite"

	"github.com/aceberg/WatchYourLAN/internal/check"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

var mu sync.Mutex

func dbExec(path, sqlStatement string) {

	mu.Lock()
	db, err := sql.Open("sqlite", path)
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
	db, _ := sql.Open("sqlite", path)
	defer db.Close()

	res, err := db.Query(sqlStatement)
	mu.Unlock()

	if err != nil {
		log.Fatal("ERROR: db_select: ", err)
	}

	dbHosts = []models.Host{}
	for res.Next() {
		var oneHost models.Host
		err = res.Scan(&oneHost.ID, &oneHost.Name, &oneHost.IP, &oneHost.Mac, &oneHost.Hw, &oneHost.Date, &oneHost.Known, &oneHost.Now)
		if err != nil {
			log.Fatal(err)
		}
		oneHost.Name = unquoteStr(oneHost.Name)
		oneHost.Hw = unquoteStr(oneHost.Hw)
		dbHosts = append(dbHosts, oneHost)
	}

	//fmt.Println("Select all:", dbHosts)
	return dbHosts
}
