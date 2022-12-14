package scan

import (
	"fmt"
	"log"

	"github.com/aceberg/WatchYourLAN/internal/db"
	"github.com/aceberg/WatchYourLAN/internal/models"
	"github.com/aceberg/WatchYourLAN/internal/notify"
)

func hostInDB(appConfig models.Conf, host models.Host, dbHosts []models.Host) bool { // Check if host is already in DB
	for _, oneHost := range dbHosts {
		if host.Mac == oneHost.Mac && (appConfig.IgnoreIP == "yes" || host.IP == oneHost.IP) {
			oneHost.IP = host.IP
			oneHost.Date = host.Date
			oneHost.Now = 1
			db.Update(appConfig.DbPath, oneHost)
			return true
		}
	}
	return false
}

func hostsCompare(appConfig models.Conf, foundHosts, dbHosts []models.Host) {
	for _, oneHost := range foundHosts {
		if !(hostInDB(appConfig, oneHost, dbHosts)) {
			oneHost.Now = 1 // Mark host online
			msg := fmt.Sprintf("UNKNOWN HOST IP: '%s', MAC: '%s', Hw: '%s'", oneHost.IP, oneHost.Mac, oneHost.Hw)
			log.Println("WARN:", msg)
			notify.Shoutrrr(msg, appConfig.ShoutURL) // Notify through Shoutrrr
			db.Insert(appConfig.DbPath, oneHost)
		}
	}
}
