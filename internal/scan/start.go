package scan

import (
	"time"

	"github.com/aceberg/WatchYourLAN/internal/db"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

var dbHosts, structHosts []models.Host
var foundHostsMap map[string]models.Host

// Start - start arp-scan goroutine
func Start(appConfig models.Conf, quit chan bool) {
	var lastDate time.Time

	db.Create(appConfig.DbPath)

	for {
		select {
		case <-quit:
			return
		default:
			nowDate := time.Now()
			plusDate := lastDate.Add(time.Duration(appConfig.Timeout) * time.Second)

			if nowDate.After(plusDate) {
				structHosts = arpScan(appConfig.Iface, appConfig.LogLevel)
				dbHosts = db.Select(appConfig.DbPath)

				toMap()
				hostsCompare(appConfig) // compare.go

				lastDate = time.Now()
			}

			time.Sleep(time.Duration(1) * time.Second)
		}
	}
}

func toMap() {
	foundHostsMap = make(map[string]models.Host)

	for _, host := range structHosts {
		foundHostsMap[host.Mac] = host
	}
}
