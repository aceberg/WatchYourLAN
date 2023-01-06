package web

import (
	"net/http"

	"github.com/aceberg/WatchYourLAN/internal/db"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig
	guiData.Hosts = AllHosts

	execTemplate(w, "index", guiData)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	AllHosts = db.Select(AppConfig.DbPath)
	http.Redirect(w, r, "/", 302)
}
