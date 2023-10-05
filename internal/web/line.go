package web

import (
	"net/http"

	"github.com/aceberg/WatchYourLAN/internal/models"
)

func lineHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	updateAllHosts()

	guiData.Config = AppConfig
	guiData.Hosts = []models.Host{}

	state := r.URL.Query().Get("state")

	for _, oneHost := range AllHosts {
		if state == "off" && oneHost.Now == 0 {
			guiData.Hosts = append(guiData.Hosts, oneHost)
		}
		if state == "on" && oneHost.Now == 1 {
			guiData.Hosts = append(guiData.Hosts, oneHost)
		}
	}

	execTemplate(w, "line", guiData)
}
