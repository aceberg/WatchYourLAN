package web

import (
	// "log/slog"
	"time"

	"github.com/aceberg/WatchYourLAN/internal/arp"
	"github.com/aceberg/WatchYourLAN/internal/db"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

func updateScan() {
	db.Create(appConfig.DBPath)

	allHosts = db.Select(appConfig.DBPath, "now")
	histHosts = db.Select(appConfig.DBPath, "history")

	quitScan = make(chan bool)
	go startScan()
}

func startScan() {
	var lastDate time.Time

	for {
		select {
		case <-quitScan:
			return
		default:
			nowDate := time.Now()
			plusDate := lastDate.Add(time.Duration(appConfig.Timeout) * time.Second)

			if nowDate.After(plusDate) {

				foundHosts := arp.Scan(appConfig.Ifaces)
				compareHosts(foundHosts)
				allHosts = db.Select(appConfig.DBPath, "now")
				histHosts = db.Select(appConfig.DBPath, "history")

				lastDate = time.Now()
			}

			time.Sleep(time.Duration(1) * time.Second)
		}
	}
}

func compareHosts(foundHosts []models.Host) {

	// Make map and Insert history
	foundHostsMap := make(map[string]models.Host)
	for _, fHost := range foundHosts {
		foundHostsMap[fHost.Mac] = fHost
		db.Insert(appConfig.DBPath, "history", fHost)
	}

	for _, aHost := range allHosts {

		fHost, exists := foundHostsMap[aHost.Mac]
		if exists && (appConfig.IgnoreIP == "yes" || aHost.IP == fHost.IP) {

			aHost.Iface = fHost.Iface
			aHost.IP = fHost.IP
			aHost.Date = fHost.Date
			aHost.Now = 1

			delete(foundHostsMap, aHost.Mac)

		} else {
			aHost.Now = 0
		}
		db.Update(appConfig.DBPath, "now", aHost)
	}

	for _, fHost := range foundHostsMap {

		// NOTIFY, LOG
		db.Insert(appConfig.DBPath, "now", fHost)
	}
}
