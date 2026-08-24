package api

import (
	"errors"
	"strconv"

	"github.com/aceberg/WatchYourLAN/internal/gdb"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

var errInvalidHostID = errors.New("invalid host id")

func getHostByID(idStr string) (models.Host, error) {

	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		return models.Host{}, errInvalidHostID
	}

	host := gdb.SelectByID(id)
	if host.ID < 1 {
		return models.Host{}, errInvalidHostID
	}

	return host, nil
}
