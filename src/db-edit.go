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
			"KNOWN"	INTEGER DEFAULT 0,
			"NOW"	INTEGER DEFAULT 0
		);`
    	db_exec(dbPath, sqlStatement)
		fmt.Println("Table created!")
    }
}

func db_insert(dbPath string, oneHost Host) {
	sqlStatement := `INSERT INTO "now" (NAME, IP, MAC, HW, DATE, KNOWN, NOW) VALUES ('%s','%s','%s','%s','%s','%d','%d');`
  	sqlStatement = fmt.Sprintf(sqlStatement, oneHost.Name, oneHost.Ip, oneHost.Mac, oneHost.Hw, oneHost.Date, oneHost.Known, oneHost.Now)
	//fmt.Println("Insert statement:", sqlStatement)
    db_exec(dbPath, sqlStatement)
}