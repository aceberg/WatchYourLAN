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
		if inString(oneHost.Name, search) || inString(oneHost.IP, search) || inString(oneHost.Mac, search) || inString(oneHost.Hw, search) || inString(oneHost.Date, search) {
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
