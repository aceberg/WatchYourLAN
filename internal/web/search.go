package web

import (
	"net/http"
	"strings"

	"github.com/aceberg/WatchYourLAN/internal/models"
)

func searchHandler(w http.ResponseWriter, r *http.Request) {
	search := r.FormValue("search")

	foundHosts := []models.Host{}
	for _, oneHost := range AllHosts {
		if inString(oneHost.Name, search) {
			foundHosts = append(foundHosts, oneHost)
		} else if inString(oneHost.IP, search) {
			foundHosts = append(foundHosts, oneHost)
		} else if inString(oneHost.Mac, search) {
			foundHosts = append(foundHosts, oneHost)
		} else if inString(oneHost.Date, search) {
			foundHosts = append(foundHosts, oneHost)
		} else if inString(oneHost.Hw, search) {
			foundHosts = append(foundHosts, oneHost)
		}
	}
	AllHosts = foundHosts

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}

func inString(str1 string, str2 string) bool {
	return strings.Contains(
		strings.ToLower(str1),
		strings.ToLower(str2),
	)
}
