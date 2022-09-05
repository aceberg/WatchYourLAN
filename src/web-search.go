package main

import (
  "net/http"
  "strings"
)

func search_hosts(w http.ResponseWriter, r *http.Request) {
	search := r.FormValue("search")

	foundHosts := []Host{}
	for _, oneHost := range AllHosts {
		if in_string(oneHost.Name, search) {
			foundHosts = append(foundHosts, oneHost)
		} else if in_string(oneHost.Ip, search) {
			foundHosts = append(foundHosts, oneHost)
		} else if in_string(oneHost.Mac, search) {
			foundHosts = append(foundHosts, oneHost)
		} else if in_string(oneHost.Date, search) {
			foundHosts = append(foundHosts, oneHost)
		} else if in_string(oneHost.Hw, search) {
			foundHosts = append(foundHosts, oneHost)
		}
	}
	AllHosts = foundHosts

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}

func in_string (str1 string, str2 string) (bool) {
	return strings.Contains(
		strings.ToLower(str1),
        strings.ToLower(str2),
	)
}