package scan

import (
	"time"

	"github.com/aceberg/WatchYourLAN/internal/db"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

// Start - start arp-scan goroutine
func Start(appConfig models.Conf) {
	var foundHosts []models.Host
	var dbHosts []models.Host
	for { // Endless
		foundHosts = arpScan(appConfig.Iface)                      // Scan interfaces
		dbHosts = db.Select(appConfig.DbPath)                      // Select everything from DB
		db.SetNow(appConfig.DbPath)                                // Mark hosts in DB as offline
		hostsCompare(appConfig.DbPath, foundHosts, dbHosts)        // Compare hosts online and in DB and add them to DB
		time.Sleep(time.Duration(appConfig.Timeout) * time.Second) // Timeout
	}
}
