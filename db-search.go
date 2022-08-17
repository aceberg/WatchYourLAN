package main

import (
	// "os"
	"fmt"
	// "log"
	// "database/sql"
	// _ "github.com/mattn/go-sqlite3"
)

func db_search(dbPath string) {
	sqlStatement := ``
    res := db_query(dbPath, sqlStatement)
	fmt.Println(res)
}