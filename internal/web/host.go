package web

import (
	"log"
	"net"
	"net/http"
	"strconv"

	"github.com/aceberg/WatchYourLAN/internal/check"
	"github.com/aceberg/WatchYourLAN/internal/db"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

func delHandler(w http.ResponseWriter, r *http.Request) {

	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	check.IfError(err)

	log.Println("INFO: delete host ID =", id)

	db.Delete(AppConfig.DbPath, id)

	http.Redirect(w, r, "/", 302)
}

func hostHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData
	var host models.Host

	guiData.Config = AppConfig

	idStr := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idStr)
	check.IfError(err)

	for _, oneHost := range AllHosts {
		if id == oneHost.ID {
			host = oneHost
			break
		}
	}

	guiData.Hosts = append(guiData.Hosts, host)

	addr, err := net.LookupAddr(host.IP)
	check.IfError(err)

	guiData.Themes = addr

	history := db.SelectHist(AppConfig.DbPath)
	for _, hist := range history {
		if hist.Host == id {
			guiData.Hist = append(guiData.Hist, hist)
		}
	}

	execTemplate(w, "host", guiData)
}
