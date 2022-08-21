package main

import (
	"os"
	"fmt"
)

func db_create() {
	if _, err := os.Stat(AppConfig.DbPath); err == nil {
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
    	db_exec(sqlStatement)
		fmt.Println("Table created!")
    }
}

func db_insert(oneHost Host) {
	sqlStatement := `INSERT INTO "now" (NAME, IP, MAC, HW, DATE, KNOWN, NOW) 
		VALUES ('%s','%s','%s','%s','%s','%d','%d');`
  	sqlStatement = fmt.Sprintf(sqlStatement, oneHost.Name, oneHost.Ip, oneHost.Mac, oneHost.Hw, oneHost.Date, oneHost.Known, oneHost.Now)
	//fmt.Println("Insert statement:", sqlStatement)
    db_exec(sqlStatement)
}

func db_update(oneHost Host) {
	sqlStatement := `UPDATE "now" set 
		NAME = '%s', IP = '%s', MAC = '%s', HW = '%s', DATE = '%s', 
		KNOWN = '%d', NOW = '%d' 
		WHERE ID = '%d';`
  	sqlStatement = fmt.Sprintf(sqlStatement, oneHost.Name, oneHost.Ip, oneHost.Mac, oneHost.Hw, oneHost.Date, oneHost.Known, oneHost.Now, oneHost.Id)
	// fmt.Println("Update statement:", sqlStatement)
    db_exec(sqlStatement)
}

func db_setnow() {
	sqlStatement := `UPDATE "now" set NOW = '0';`
	// fmt.Println("Set now 0 statement:", sqlStatement)
    db_exec(sqlStatement)
}