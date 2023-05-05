package scan

import (
	"time"

	"github.com/aceberg/WatchYourLAN/internal/db"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

// Start - start arp-scan goroutine
func Start(appConfig models.Conf, quit chan bool) {
	var foundHosts []models.Host
	var dbHosts []models.Host
	var lastDate time.Time

	for {
		select {
		case <-quit:
			return
		default:
			nowDate := time.Now()
			plusDate := lastDate.Add(time.Duration(appConfig.Timeout) * time.Second)

			if nowDate.After(plusDate) {
				foundHosts = arpScan(appConfig.Iface, appConfig.LogLevel)
				dbHosts = db.Select(appConfig.DbPath)
				db.SetNow(appConfig.DbPath)
				hostsCompare(appConfig, foundHosts, dbHosts)

				lastDate = time.Now()
			}

			time.Sleep(time.Duration(1) * time.Second)
		}
	}
}
