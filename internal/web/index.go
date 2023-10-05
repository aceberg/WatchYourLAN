package web

import (
	"net/http"

	"github.com/aceberg/WatchYourLAN/internal/db"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	AllHosts = db.Select(AppConfig.DbPath)

	guiData.Config = AppConfig
	guiData.Hosts = AllHosts

	execTemplate(w, "index", guiData)
}
