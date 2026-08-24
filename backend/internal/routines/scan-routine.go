package routines

import (
	"log/slog"
	"time"

	"github.com/aceberg/WatchYourLAN/internal/arp"
	"github.com/aceberg/WatchYourLAN/internal/check"
	"github.com/aceberg/WatchYourLAN/internal/conf"
	"github.com/aceberg/WatchYourLAN/internal/gdb"
	"github.com/aceberg/WatchYourLAN/internal/influx"
	"github.com/aceberg/WatchYourLAN/internal/models"
	"github.com/aceberg/WatchYourLAN/internal/notify"
	"github.com/aceberg/WatchYourLAN/internal/prometheus"
)

func startScan(quit chan bool) {
	var lastDate, nowDate, plusDate time.Time
	var foundHosts []models.Host

	for {
		select {
		case <-quit:
			return
		default:
			config := conf.GetAppConfig()
			nowDate = time.Now()
			plusDate = lastDate.Add(time.Duration(config.Timeout) * time.Second)

			if nowDate.After(plusDate) {

				var scanOK bool
				foundHosts, scanOK = arp.Scan(config.Ifaces, config.ArpArgs, config.ArpStrs)
				if !processScanResult(foundHosts, scanOK) {
					lastDate = time.Now()
					continue
				}

				lastDate = time.Now()
			}

			time.Sleep(time.Duration(1) * time.Minute)
		}
	}
}

func processScanResult(foundHosts []models.Host, scanOK bool) bool {
	if !scanOK {
		slog.Warn("Skipping host state update because ARP scan failed")
		return false
	}

	foundHostsMap := make(map[string]models.Host)
	for _, fHost := range foundHosts {
		foundHostsMap[fHost.Mac] = fHost
	}

	compareHosts(foundHostsMap)
	return true
}

func compareHosts(foundHostsMap map[string]models.Host) {
	config := conf.GetAppConfig()

	allHosts, ok := gdb.Select("now")
	if !ok {
		return
	}

	for _, aHost := range allHosts {
		previousNow := aHost.Now

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
		gdb.Update("now", aHost)

		if exists && previousNow == 0 {
			gdb.RecordHostEvent(aHost, models.EventOnline, "", "")
		}
		if !exists && previousNow == 1 {
			gdb.RecordHostEvent(aHost, models.EventOffline, "", "")
		}

		aHost.ID = 0
		aHost.Date = time.Now().Format("2006-01-02 15:04:05")
		gdb.Update("history", aHost)

		if config.InfluxEnable {
			influx.Add(config, aHost)
		}
		if config.PrometheusEnable {
			prometheus.Add(aHost)
		}
	}

	for _, fHost := range foundHostsMap {

		fHost.Name, fHost.DNS = check.DNS(fHost)
		notify.Unknown(fHost) // Log and Shoutrrr

		gdb.Update("now", fHost)
		hosts := gdb.SelectByMAC("now", fHost.Mac)
		if len(hosts) > 0 {
			gdb.RecordHostEvent(hosts[0], models.EventDiscovered, "", "")
		}
	}
}
