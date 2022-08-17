package main

import (
	"os"
	"fmt"
)

func db_create(dbPath string) {
	if _, err := os.Stat(dbPath); err == nil {
        fmt.Println("DB exists")
    } else {
		sqlStatement := `CREATE TABLE "now" (
			"ID"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
			"NAME"	TEXT NOT NULL,
			"IP"	TEXT,
			"MAC"	TEXT,
			"HW"	TEXT,
			"DATE"	TEXT,
			"KNOWN"	INTEGER DEFAULT 0
		);`
    	db_exec(dbPath, sqlStatement)
		fmt.Println("Table created!")
    }
}

func db_insert(dbPath string, oneHost Host) {
	sqlStatement := `INSERT INTO "now" (NAME, IP, MAC, HW, DATE, KNOWN) VALUES ('','%s','%s','%s','%s', '%d');`
  	sqlStatement = fmt.Sprintf(sqlStatement, oneHost.Ip, oneHost.Mac, oneHost.Hw, oneHost.Date, 1)
	fmt.Println("Insert statement:", sqlStatement)
    db_exec(dbPath, sqlStatement)
}