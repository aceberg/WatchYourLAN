package gdb

import (
	"github.com/aceberg/WatchYourLAN/internal/check"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

// Select - get all hosts
func Select(table string) (dbHosts []models.Host) {

	tab := db.Table(table)
	tab.Find(&dbHosts)

	return dbHosts
}

// SelectByID - get host by ID
func SelectByID(id int) (host models.Host) {

	tab := db.Table("now")
	tab.First(&host, id)

	return host
}

// SelectByMAC - get all hosts by MAC
func SelectByMAC(mac string) (hosts []models.Host) {

	tab := db.Table("history")
	tab.Where("MAC = ?", mac).Find(&hosts)

	return hosts
}

// Update - update or create host
func Update(table string, oneHost models.Host) {

	tab := db.Table(table)
	result := tab.Save(&oneHost)
	check.IfError(result.Error)
}

// Delete - delete host from DB
func Delete(table string, id int) {

	tab := db.Table(table)
	result := tab.Delete(&models.Host{}, id)
	check.IfError(result.Error)
}

// DeleteOldHistory - delete a list of hosts from History
func DeleteOldHistory(date string) int64 {

	tab := db.Table("history")
	result := tab.Where("DATE < ?", date).Delete(&models.Host{})
	check.IfError(result.Error)

	return result.RowsAffected
}

// Clear - delete all hosts from table
func Clear(table string) {

	tab := db.Table(table)
	result := tab.Where("1 = 1").Delete(&models.Host{})
	check.IfError(result.Error)
}
