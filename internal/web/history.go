package web

import (
	"net/http"

	"github.com/aceberg/WatchYourLAN/internal/db"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

func historyHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig
	guiData.Hist = db.SelectHist(AppConfig.DbPath)

	execTemplate(w, "history", guiData)
}
