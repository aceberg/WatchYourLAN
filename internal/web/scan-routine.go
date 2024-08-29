package web

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/aceberg/WatchYourLAN/internal/arp"
	"github.com/aceberg/WatchYourLAN/internal/db"
	"github.com/aceberg/WatchYourLAN/internal/influx"
	"github.com/aceberg/WatchYourLAN/internal/models"
	"github.com/aceberg/WatchYourLAN/internal/notify"
)

func startScan(quit chan bool) {
	var lastDate, nowDate, plusDate time.Time
	var foundHosts []models.Host

	for {
		select {
		case <-quit:
			return
		default:
			nowDate = time.Now()
			plusDate = lastDate.Add(time.Duration(appConfig.Timeout) * time.Second)

			if nowDate.After(plusDate) {

				foundHosts = arp.Scan(appConfig.Ifaces, appConfig.ArpArgs)
				compareHosts(foundHosts)
				allHosts = db.Select("now")

				lastDate = time.Now()
			}

			time.Sleep(time.Duration(1) * time.Second)
		}
	}
}

func compareHosts(foundHosts []models.Host) {

	// Make map of found hosts
	foundHostsMap := make(map[string]models.Host)
	for _, fHost := range foundHosts {
		foundHostsMap[fHost.Mac] = fHost
	}

	for _, aHost := range allHosts {

		fHost, exists := foundHostsMap[aHost.Mac]
		if exists {

			aHost.Iface = fHost.Iface
			aHost.IP = fHost.IP
			aHost.Date = fHost.Date
			aHost.Now = 1

			delete(foundHostsMap, aHost.Mac)

		} else {
			aHost.Now = 0
		}
		db.Update("now", aHost)

		aHost.Date = time.Now().Format("2006-01-02 15:04:05")
		histHosts = append(histHosts, aHost)
		if appConfig.HistInDB {
			db.Insert("history", aHost)
		}
		if appConfig.InfluxEnable {
			influx.Add(appConfig, aHost)
		}
	}

	for _, fHost := range foundHostsMap {

		msg := fmt.Sprintf("Unknown host IP: '%s', MAC: '%s', Hw: '%s', Iface: '%s'", fHost.IP, fHost.Mac, fHost.Hw, fHost.Iface)
		slog.Warn(msg)
		notify.Shout(msg, appConfig.ShoutURL) // Notify through Shoutrrr

		fHost.Name, fHost.DNS = updateDNS(fHost)
		db.Insert("now", fHost)
	}
}
