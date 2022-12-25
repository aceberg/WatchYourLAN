package db

import (
	//"fmt"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func dbExec(path, sqlStatement string) {
	db, _ := sql.Open("sqlite3", path)
	defer db.Close()

	_, err := db.Exec(sqlStatement)
	if err != nil {
		log.Fatal("ERROR: db_exec: ", err)
	}
}
