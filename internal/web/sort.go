package web

import (
	"bytes"
	"net"
	"net/http"
	"reflect"
	"sort"
	"strconv"

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
		sortByField("asc", "Name")
	case "name-down":
		sortByField("desc", "Name")
	case "ip-up":
		sortByIPs("asc")
	case "ip-down":
		sortByIPs("desc")
	case "mac-up":
		sortByField("asc", "Mac")
	case "mac-down":
		sortByField("desc", "Mac")
	case "hw-up":
		sortByField("asc", "Hw")
	case "hw-down":
		sortByField("desc", "Hw")
	case "date-up":
		sortByField("asc", "Date")
	case "date-down":
		sortByField("desc", "Date")
	case "known-up":
		sortByField("asc", "Known")
	case "known-down":
		sortByField("desc", "Known")
	default:
		updateAllHosts()
	}

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}

func sortByField(method string, field string) {
	type fieldHost struct {
		Host models.Host
		F    string
	}
	toSort := []fieldHost{}
	var oneSort fieldHost

	for _, host := range AllHosts {
		oneSort.Host = host
		r := reflect.ValueOf(&host)
		f := reflect.Indirect(r).FieldByName(field)
		if field == "Known" {
			oneSort.F = strconv.FormatUint(f.Uint(), 10)
		} else {
			oneSort.F = f.String()
		}
		toSort = append(toSort, oneSort)
	}

	switch method {
	case "asc":
		sort.Slice(toSort, func(i, j int) bool {
			return toSort[i].F < toSort[j].F
		})
	case "desc":
		sort.Slice(toSort, func(i, j int) bool {
			return toSort[i].F > toSort[j].F
		})
	}

	AllHosts = []models.Host{}
	for _, oneSort := range toSort {
		AllHosts = append(AllHosts, oneSort.Host)
	}
}
