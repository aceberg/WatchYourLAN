package main

import (
  "net/http"
  "sort"
  "bytes"
  "net"
)

func sort_by_ips(method string) {
	type ipHost struct {
		Host Host
		Ip net.IP
	}
	toSort := []ipHost{}
	var oneSort ipHost

	for _, oneHost := range AllHosts {
		oneSort.Host = oneHost
		oneSort.Ip = net.ParseIP(oneHost.Ip)
		toSort = append(toSort, oneSort)
	}

	switch method {
	case "asc":
		sort.Slice(toSort, func(i, j int) bool {
			return bytes.Compare(toSort[i].Ip, toSort[j].Ip) < 0
		})
	case "desc":
		sort.Slice(toSort, func(i, j int) bool {
			return bytes.Compare(toSort[i].Ip, toSort[j].Ip) > 0
		})
	}

	AllHosts = []Host{}
	for _, oneSort := range toSort {
		AllHosts = append(AllHosts, oneSort.Host)
	}
}

func sort_hosts(w http.ResponseWriter, r *http.Request) {

	sort_method := r.FormValue("sort_method")

	switch sort_method {
	case "name-up":
		sort.SliceStable(AllHosts, func(i, j int) bool {
			return AllHosts[i].Name < AllHosts[j].Name
		})
	case "name-down":
		sort.SliceStable(AllHosts, func(i, j int) bool {
			return AllHosts[i].Name > AllHosts[j].Name
		})
	case "ip-up":
		sort_by_ips("asc")
	case "ip-down":
		sort_by_ips("desc")
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
		AllHosts = db_select()
	}

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}