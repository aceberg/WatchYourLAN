package db

import (
	"fmt"

	"github.com/aceberg/WatchYourLAN/internal/models"
)

// InsertHist - insert history into table
func InsertHist(path string, hist models.History) {
	hist.Name = quoteStr(hist.Name)
	hist.Hw = quoteStr(hist.Hw)
	sqlStatement := `INSERT INTO "history" (HOST, NAME, IP, MAC, HW, DATE, KNOWN, STATE) 
		VALUES ('%d','%s','%s','%s','%s','%s','%d','%d');`
	sqlStatement = fmt.Sprintf(sqlStatement, hist.Host, hist.Name, hist.IP, hist.Mac, hist.Hw, hist.Date, hist.Known, hist.State)
	//fmt.Println("Insert statement:", sqlStatement)
	dbExec(path, sqlStatement)
}

// DeleteHist - delete history from DB
func DeleteHist(path string, id int) {
	sqlStatement := `DELETE FROM "history" WHERE ID='%d';`
	sqlStatement = fmt.Sprintf(sqlStatement, id)
	dbExec(path, sqlStatement)
}

// ClearHist - delete all history from table
func ClearHist(path string) {
	sqlStatement := `DELETE FROM "history";`
	dbExec(path, sqlStatement)
}
