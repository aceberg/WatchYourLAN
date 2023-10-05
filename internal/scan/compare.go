package scan

import (
	"fmt"
	"log"

	"github.com/aceberg/WatchYourLAN/internal/db"
	"github.com/aceberg/WatchYourLAN/internal/models"
	"github.com/aceberg/WatchYourLAN/internal/notify"
)

func hostsCompare(appConfig models.Conf) {

	for _, oneHost := range dbHosts {

		host, exists := foundHostsMap[oneHost.Mac]

		if exists && (appConfig.IgnoreIP == "yes" || host.IP == oneHost.IP) {
			oneHost.Date = host.Date
			oneHost.Now = 1

			delete(foundHostsMap, oneHost.Mac)

			db.Update(appConfig.DbPath, oneHost)

		} else if oneHost.Now == 1 {
			oneHost.Now = 0
			db.Update(appConfig.DbPath, oneHost)
		}
	}

	for _, oneHost := range foundHostsMap {

		msg := fmt.Sprintf("UNKNOWN HOST IP: '%s', MAC: '%s', Hw: '%s'", oneHost.IP, oneHost.Mac, oneHost.Hw)
		log.Println("WARN:", msg)
		notify.Shoutrrr(msg, appConfig.ShoutURL) // Notify through Shoutrrr

		db.Insert(appConfig.DbPath, oneHost)
	}
}
