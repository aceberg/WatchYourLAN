package routines

import (
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
			nowDate = time.Now()
			plusDate = lastDate.Add(time.Duration(conf.AppConfig.Timeout) * time.Second)

			if nowDate.After(plusDate) {

				foundHosts = arp.Scan(conf.AppConfig.Ifaces, conf.AppConfig.ArpArgs, conf.AppConfig.ArpStrs)

				// Make map of found hosts
				foundHostsMap := make(map[string]models.Host)
				for _, fHost := range foundHosts {
					foundHostsMap[fHost.Mac] = fHost
				}

				compareHosts(foundHostsMap)

				lastDate = time.Now()
			}

			time.Sleep(time.Duration(1) * time.Minute)
		}
	}
}

func compareHosts(foundHostsMap map[string]models.Host) {

	allHosts, ok := gdb.Select("now")
	if !ok {
		return
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
		gdb.Update("now", aHost)

		aHost.ID = 0
		aHost.Date = time.Now().Format("2006-01-02 15:04:05")
		gdb.Update("history", aHost)

		if conf.AppConfig.InfluxEnable {
			influx.Add(conf.AppConfig, aHost)
		}
		if conf.AppConfig.PrometheusEnable {
			prometheus.Add(aHost)
		}
	}

	for _, fHost := range foundHostsMap {

		fHost.Name, fHost.DNS = check.DNS(fHost)
		notify.Unknown(fHost) // Log and Shoutrrr

		gdb.Update("now", fHost)
	}
}
