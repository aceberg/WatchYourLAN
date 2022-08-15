package main

import (
	"os"
	"fmt"
	"log"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func db_create(dbPath string) {
	if _, err := os.Stat(dbPath); err == nil {
        fmt.Println("DB exists")
    } else {
		sqlStatement := `CREATE TABLE "now" (
			"ID"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
			"NAME"	TEXT,
			"IP"	TEXT,
			"MAC"	TEXT,
			"HW"	TEXT,
			"DATE"	TEXT,
			"KNOWN"	INTEGER DEFAULT 0
		);`
    	db_exec(dbPath, sqlStatement)
    }
}

func db_exec (dbPath string, sqlStatement string) {
	db, _ := sql.Open("sqlite3", dbPath)
	defer db.Close()
  
	_, err := db.Exec(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}
}