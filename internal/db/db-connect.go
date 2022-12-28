package db

import (
	"database/sql"
	"log"

	// Import sqlite module
	_ "modernc.org/sqlite"

	"github.com/aceberg/WatchYourLAN/internal/check"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

func dbExec(path, sqlStatement string) {
	db, err := sql.Open("sqlite", path)
	check.IfError(err)
	defer db.Close()

	_, err = db.Exec(sqlStatement)
	check.IfError(err)
}

// Select - select all hosts
func Select(path string) (dbHosts []models.Host) {
	db, _ := sql.Open("sqlite", path)
	defer db.Close()

	sqlStatement := `SELECT * FROM "now" ORDER BY DATE DESC`

	res, err := db.Query(sqlStatement)
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
