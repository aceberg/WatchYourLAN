package scan

import (
	"time"

	"github.com/aceberg/WatchYourLAN/internal/db"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

var appConfig models.Conf
var dbHosts, structHosts []models.Host
var foundHostsMap map[string]models.Host

// Start - start arp-scan goroutine
func Start(config models.Conf, quit chan bool) {
	var lastDate time.Time

	appConfig = config
	db.Create(appConfig.DbPath)
	dbHosts = db.Select(appConfig.DbPath)

	for {
		select {
		case <-quit:
			return
		default:
			nowDate := time.Now()
			plusDate := lastDate.Add(time.Duration(appConfig.Timeout) * time.Second)

			if nowDate.After(plusDate) {
				structHosts = arpScan(appConfig.Iface, appConfig.LogLevel) // arpscan.go
				dbHosts = db.Select(appConfig.DbPath)

				toMap()
				hostsCompare() // compare.go

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
