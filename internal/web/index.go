package web

import (
	"net/http"

	"github.com/aceberg/WatchYourLAN/internal/models"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	updateAllHosts()

	guiData.Config = AppConfig
	guiData.Hosts = AllHosts

	execTemplate(w, "index", guiData)
}
