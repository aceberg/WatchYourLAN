package db

import (
	"fmt"
	"log"
	"os"

	"github.com/aceberg/WatchYourLAN/internal/models"
)

// Create - create DB if not exists
func Create(path string) {
	if _, err := os.Stat(path); err == nil {
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
		dbExec(path, sqlStatement)
		log.Println("INFO: Table created!")
	}
}

// SetNow - mark all hosts as offline
func SetNow(path string) {
	sqlStatement := `UPDATE "now" set NOW = '0';`
	dbExec(path, sqlStatement)
}

// Insert - insert host into table
func Insert(path string, oneHost models.Host) {
	oneHost.Name = quoteStr(oneHost.Name)
	oneHost.Hw = quoteStr(oneHost.Hw)
	sqlStatement := `INSERT INTO "now" (NAME, IP, MAC, HW, DATE, KNOWN, NOW) 
		VALUES ('%s','%s','%s','%s','%s','%d','%d');`
	sqlStatement = fmt.Sprintf(sqlStatement, oneHost.Name, oneHost.IP, oneHost.Mac, oneHost.Hw, oneHost.Date, oneHost.Known, oneHost.Now)
	//fmt.Println("Insert statement:", sqlStatement)
	dbExec(path, sqlStatement)
}

// Update - update host
func Update(path string, oneHost models.Host) {
	oneHost.Name = quoteStr(oneHost.Name)
	oneHost.Hw = quoteStr(oneHost.Hw)
	sqlStatement := `UPDATE "now" set 
		NAME = '%s', IP = '%s', MAC = '%s', HW = '%s', DATE = '%s', 
		KNOWN = '%d', NOW = '%d' 
		WHERE ID = '%d';`
	sqlStatement = fmt.Sprintf(sqlStatement, oneHost.Name, oneHost.IP, oneHost.Mac, oneHost.Hw, oneHost.Date, oneHost.Known, oneHost.Now, oneHost.ID)
	//fmt.Println("Update statement:", sqlStatement)
	dbExec(path, sqlStatement)
}
