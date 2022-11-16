package db

import (
	"os"
	"log"
	"fmt"
)

func db_create() {
	if _, err := os.Stat(AppConfig.DbPath); err == nil {
        log.Println("INFO: DB exists")
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
		log.Println("INFO: Table created!")
    }
}

func db_insert(oneHost Host) {
	oneHost.Name = quote_str(oneHost.Name)
	oneHost.Hw = quote_str(oneHost.Hw)
	sqlStatement := `INSERT INTO "now" (NAME, IP, MAC, HW, DATE, KNOWN, NOW) 
		VALUES ('%s','%s','%s','%s','%s','%d','%d');`
  	sqlStatement = fmt.Sprintf(sqlStatement, oneHost.Name, oneHost.Ip, oneHost.Mac, oneHost.Hw, oneHost.Date, oneHost.Known, oneHost.Now)
	//fmt.Println("Insert statement:", sqlStatement)
    db_exec(sqlStatement)
}

func db_update(oneHost Host) {
	oneHost.Name = quote_str(oneHost.Name)
	oneHost.Hw = quote_str(oneHost.Hw)
	sqlStatement := `UPDATE "now" set 
		NAME = '%s', IP = '%s', MAC = '%s', HW = '%s', DATE = '%s', 
		KNOWN = '%d', NOW = '%d' 
		WHERE ID = '%d';`
  	sqlStatement = fmt.Sprintf(sqlStatement, oneHost.Name, oneHost.Ip, oneHost.Mac, oneHost.Hw, oneHost.Date, oneHost.Known, oneHost.Now, oneHost.Id)
	//fmt.Println("Update statement:", sqlStatement)
    db_exec(sqlStatement)
}

func db_setnow() {
	sqlStatement := `UPDATE "now" set NOW = '0';`
    db_exec(sqlStatement)
}