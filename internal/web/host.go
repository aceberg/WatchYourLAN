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

	db.Delete(AppConfig.DbPath, uint16(id))

	http.Redirect(w, r, "/home/", 302)
}

func hostHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData
	var host models.Host

	guiData.Config = AppConfig
	guiData.Icon = Icon

	idStr := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idStr)
	check.IfError(err)

	id16 := uint16(id)

	for _, oneHost := range AllHosts {
		if id16 == oneHost.ID {
			host = oneHost
			break
		}
	}

	guiData.Hosts = append(guiData.Hosts, host)

	addr, err := net.LookupAddr(host.IP)
	check.IfError(err)

	guiData.Themes = addr

	execTemplate(w, "host", guiData)
}
