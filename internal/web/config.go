package web

import (
	"log"
	"net/http"
	"strconv"

	"github.com/aceberg/WatchYourLAN/internal/check"
	"github.com/aceberg/WatchYourLAN/internal/conf"
	"github.com/aceberg/WatchYourLAN/internal/db"
	"github.com/aceberg/WatchYourLAN/internal/models"
	"github.com/aceberg/WatchYourLAN/internal/scan"
)

func configHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig
	guiData.Icon = Icon

	guiData.Themes = []string{"cerulean", "cosmo", "cyborg", "darkly", "flatly", "journal", "litera", "lumen", "lux", "materia", "minty", "morph", "pulse", "quartz", "sandstone", "simplex", "sketchy", "slate", "solar", "spacelab", "superhero", "united", "vapor", "yeti", "zephyr"}

	execTemplate(w, "config", guiData)
}

func saveConfigHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	AppConfig.Iface = r.FormValue("iface")
	AppConfig.DbPath = r.FormValue("dbpath")
	AppConfig.GuiIP = r.FormValue("host")
	AppConfig.GuiPort = r.FormValue("port")
	AppConfig.ShoutURL = r.FormValue("shout")
	AppConfig.Theme = r.FormValue("theme")
	AppConfig.IgnoreIP = r.FormValue("ignoreip")

	timeout := r.FormValue("timeout")
	AppConfig.Timeout, err = strconv.Atoi(timeout)
	check.IfError(err)

	close(QuitScan)

	conf.Write(ConfigPath, AppConfig)

	QuitScan = make(chan bool)
	go scan.Start(AppConfig, QuitScan)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}

func clearHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("INFO: delting all hosts from DB")

	db.Clear(AppConfig.DbPath)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}
