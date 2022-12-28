package web

import (
	"bytes"
	"net"
	"net/http"
	"sort"

	"github.com/aceberg/WatchYourLAN/internal/db"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

func sortByIPs(method string) {
	type ipHost struct {
		Host models.Host
		IP   net.IP
	}
	toSort := []ipHost{}
	var oneSort ipHost

	for _, oneHost := range AllHosts {
		oneSort.Host = oneHost
		oneSort.IP = net.ParseIP(oneHost.IP)
		toSort = append(toSort, oneSort)
	}

	switch method {
	case "asc":
		sort.Slice(toSort, func(i, j int) bool {
			return bytes.Compare(toSort[i].IP, toSort[j].IP) < 0
		})
	case "desc":
		sort.Slice(toSort, func(i, j int) bool {
			return bytes.Compare(toSort[i].IP, toSort[j].IP) > 0
		})
	}

	AllHosts = []models.Host{}
	for _, oneSort := range toSort {
		AllHosts = append(AllHosts, oneSort.Host)
	}
}

func sortHandler(w http.ResponseWriter, r *http.Request) {

	sortMethod := r.FormValue("sort_method")

	switch sortMethod {
	case "name-up":
		sort.SliceStable(AllHosts, func(i, j int) bool {
			return AllHosts[i].Name < AllHosts[j].Name
		})
	case "name-down":
		sort.SliceStable(AllHosts, func(i, j int) bool {
			return AllHosts[i].Name > AllHosts[j].Name
		})
	case "ip-up":
		sortByIPs("asc")
	case "ip-down":
		sortByIPs("desc")
	case "date-up":
		sort.SliceStable(AllHosts, func(i, j int) bool {
			return AllHosts[i].Date < AllHosts[j].Date
		})
	case "date-down":
		sort.SliceStable(AllHosts, func(i, j int) bool {
			return AllHosts[i].Date > AllHosts[j].Date
		})
	case "known-up":
		sort.SliceStable(AllHosts, func(i, j int) bool {
			return AllHosts[i].Known < AllHosts[j].Known
		})
	case "known-down":
		sort.SliceStable(AllHosts, func(i, j int) bool {
			return AllHosts[i].Known > AllHosts[j].Known
		})
	default:
		AllHosts = db.Select(AppConfig.DbPath)
	}

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}
