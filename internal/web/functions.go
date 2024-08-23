package web

import (
	"github.com/aceberg/WatchYourLAN/internal/models"
)

func getHostByID(id int, hosts []models.Host) (oneHost models.Host) {

	for _, host := range hosts {
		if host.ID == id {
			oneHost = host
			break
		}
	}

	return oneHost
}
