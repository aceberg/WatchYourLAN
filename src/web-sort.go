package main

import (
  "net/http"
  "sort"
)

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
		sort.SliceStable(AllHosts, func(i, j int) bool {
			return AllHosts[i].Ip < AllHosts[j].Ip
		})
	case "ip-down":
		sort.SliceStable(AllHosts, func(i, j int) bool {
			return AllHosts[i].Ip > AllHosts[j].Ip
		})
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

	http.Redirect(w, r, "/", http.StatusSeeOther)
}