package db

import (
	"log"
	"os"
	// "fmt"
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
