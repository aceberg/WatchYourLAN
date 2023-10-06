package db

import (
	"fmt"

	"github.com/aceberg/WatchYourLAN/internal/models"
)

// Create - create DB if not exists
func Create(path string) {

	sqlStatement := `CREATE TABLE IF NOT EXISTS "now" (
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

	sqlStatement = `CREATE TABLE IF NOT EXISTS "history" (
		"ID"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
		"HOST"	INTEGER DEFAULT 0,
		"NAME"	TEXT NOT NULL,
		"IP"	TEXT,
		"MAC"	TEXT,
		"HW"	TEXT,
		"DATE"	TEXT,
		"KNOWN"	INTEGER DEFAULT 0,
		"STATE"	INTEGER DEFAULT 0
	);`
	dbExec(path, sqlStatement)
}

// Insert - insert host into table
func Insert(path string, oneHost models.Host) {
	oneHost.Name = quoteStr(oneHost.Name)
	oneHost.Hw = quoteStr(oneHost.Hw)
	sqlStatement := `INSERT INTO "now" (NAME, IP, MAC, HW, DATE, KNOWN, NOW) 
		VALUES ('%s','%s','%s','%s','%s','%d','%d');`
	sqlStatement = fmt.Sprintf(sqlStatement, oneHost.Name, oneHost.IP, oneHost.Mac, oneHost.Hw, oneHost.Date, oneHost.Known, oneHost.Now)

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

	dbExec(path, sqlStatement)
}

// Delete - delete host from DB
func Delete(path string, id int) {
	sqlStatement := `DELETE FROM "now" WHERE ID='%d';`
	sqlStatement = fmt.Sprintf(sqlStatement, id)
	dbExec(path, sqlStatement)
}

// Clear - delete all hosts from table
func Clear(path string) {
	sqlStatement := `DELETE FROM "now";`
	dbExec(path, sqlStatement)
}
