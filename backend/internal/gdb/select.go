package gdb

import (
	"github.com/aceberg/WatchYourLAN/internal/check"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

// Select - get all hosts
func Select(table string) (dbHosts []models.Host, ok bool) {

	tab := db.Table(table)
	err := tab.Find(&dbHosts).Error

	return dbHosts, !check.IfError(err)
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
	tab.Where("\"MAC\" = ?", mac).Find(&hosts)

	return hosts
}

// SelectByDate - get all hosts by MAC and DATE
func SelectByDate(mac, date string) (hosts []models.Host) {

	tab := db.Table("history")
	tab.
		Where("\"MAC\" = ?", mac).
		Where("\"DATE\" LIKE ?", date+"%").
		Find(&hosts)

	return hosts
}

// SelectLatest - get latest hosts by MAC
func SelectLatest(mac string, number int) (hosts []models.Host) {

	tab := db.Table("history")
	tab.
		Where("\"MAC\" = ?", mac).
		Order("\"DATE\" DESC").
		Limit(number).
		Find(&hosts)

	return hosts
}
