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

// DeleteList - delete a list of hosts from History
func DeleteList(ids []int) {

	tab := db.Table("history")
	result := tab.Where("ID IN ?", ids).Delete(&models.Host{})
	check.IfError(result.Error)
}

// Clear - delete all hosts from table
func Clear(table string) {

	tab := db.Table(table)
	result := tab.Where("1 = 1").Delete(&models.Host{})
	check.IfError(result.Error)
}
