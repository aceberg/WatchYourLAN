package db

import (
	"fmt"

	"github.com/aceberg/WatchYourLAN/internal/models"
)

// Create - create DB if not exists
func Create() {

	sqlStatement := `CREATE TABLE IF NOT EXISTS "now" (
		"ID"	    ` + currentDB.PrimaryKey + `,
		"NAME"	    TEXT NOT NULL,
		"DNS"	    TEXT NOT NULL,
		"IFACE"	    TEXT,
		"IP"	    TEXT,
		"MAC"	    TEXT,
		"HW"	    TEXT,
		"DATE"	    TEXT,
		"FIRSTSEEN" TEXT,
		"KNOWN"	    INTEGER DEFAULT 0,
		"NOW"	    INTEGER DEFAULT 0
	);`
	dbExec(sqlStatement)

	sqlStatement = `CREATE TABLE IF NOT EXISTS "history" (
		"ID"	    ` + currentDB.PrimaryKey + `,
		"NAME"	    TEXT NOT NULL,
		"DNS"	    TEXT NOT NULL,
		"IFACE"	    TEXT,
		"IP"	    TEXT,
		"MAC"	    TEXT,
		"HW"	    TEXT,
		"DATE"	    TEXT,
		"FIRSTSEEN" TEXT,
		"KNOWN"	    INTEGER DEFAULT 0,
		"NOW"	    INTEGER DEFAULT 0
	);`
	dbExec(sqlStatement)
}

// Insert - insert host into table
func Insert(table string, oneHost models.Host) {
	oneHost.Name = quoteStr(oneHost.Name)
	oneHost.Hw = quoteStr(oneHost.Hw)
	sqlStatement := `INSERT INTO %s ("NAME", "DNS", "IFACE", "IP", "MAC", "HW", "DATE", "FIRSTSEEN", "KNOWN", "NOW") VALUES ('%s','%s','%s','%s','%s','%s','%s','%s','%d','%d');`
	sqlStatement = fmt.Sprintf(sqlStatement, table, oneHost.Name, oneHost.DNS, oneHost.Iface, oneHost.IP, oneHost.Mac, oneHost.Hw, oneHost.Date, oneHost.FirstSeen, oneHost.Known, oneHost.Now)

	dbExec(sqlStatement)
}

// Update - update host
func Update(table string, oneHost models.Host) {
	oneHost.Name = quoteStr(oneHost.Name)
	oneHost.Hw = quoteStr(oneHost.Hw)
	sqlStatement := `UPDATE %s set 
		"NAME" = '%s', "DNS" = '%s', "IFACE" = '%s', "IP" = '%s', "MAC" = '%s', "HW" = '%s', "DATE" = '%s', "KNOWN" = '%d', "NOW" = '%d' 
		WHERE "ID" = '%d';`
	sqlStatement = fmt.Sprintf(sqlStatement, table, oneHost.Name, oneHost.DNS, oneHost.Iface, oneHost.IP, oneHost.Mac, oneHost.Hw, oneHost.Date, oneHost.Known, oneHost.Now, oneHost.ID)

	dbExec(sqlStatement)
}

// Delete - delete host from DB
func Delete(table string, id int) {
	sqlStatement := `DELETE FROM %s WHERE "ID"='%d';`
	sqlStatement = fmt.Sprintf(sqlStatement, table, id)
	dbExec(sqlStatement)
}

// DeleteList - delete a list of hosts from History
func DeleteList(ids []int) {
	if len(ids) > 0 {
		idString := ""

		for _, id := range ids {
			idString = idString + fmt.Sprintf("%d, ", id)
		}
		idString = idString[:len(idString)-2]

		sqlStatement := `DELETE FROM history WHERE "ID" IN (%s);`
		sqlStatement = fmt.Sprintf(sqlStatement, idString)

		dbExec(sqlStatement)
	}
}

// Clear - delete all hosts from table
func Clear(table string) {
	sqlStatement := `DELETE FROM %s;`
	sqlStatement = fmt.Sprintf(sqlStatement, table)
	dbExec(sqlStatement)
}
