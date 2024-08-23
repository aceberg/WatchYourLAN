package web

import (
	"strconv"

	"github.com/aceberg/WatchYourLAN/internal/models"
)

func getHostByID(idStr string, hosts []models.Host) (oneHost models.Host) {

	id, _ := strconv.Atoi(idStr)

	for _, host := range hosts {
		if host.ID == id {
			oneHost = host
			break
		}
	}

	return oneHost
}
