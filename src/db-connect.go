package main

import (
	//"fmt"
	"log"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func db_exec(sqlStatement string) {
	db, _ := sql.Open("sqlite3", AppConfig.DbPath)
	defer db.Close()
  
	_, err := db.Exec(sqlStatement)
	if err != nil {
		log.Fatal("ERROR: db_exec: ", err)
	}
}

func db_select() (dbHosts []Host) {
	db, _ := sql.Open("sqlite3", AppConfig.DbPath)
	defer db.Close()

	sqlStatement := `SELECT * FROM "now" ORDER BY DATE DESC`

	res, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatal("ERROR: db_select: ", err)
	}

	dbHosts = []Host{}
  	for res.Next() {
    	var oneHost Host
    	err = res.Scan(&oneHost.Id, &oneHost.Name, &oneHost.Ip, &oneHost.Mac, &oneHost.Hw, &oneHost.Date, &oneHost.Known, &oneHost.Now)
    	if err != nil {
			log.Fatal(err)
    	}
		oneHost.Name = unquote_str(oneHost.Name)
    	dbHosts = append(dbHosts, oneHost)
  	}

	//fmt.Println("Select all:", dbHosts)
	return dbHosts
}