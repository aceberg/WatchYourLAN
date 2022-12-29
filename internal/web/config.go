package web

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/aceberg/WatchYourLAN/internal/check"
	"github.com/aceberg/WatchYourLAN/internal/conf"
	"github.com/aceberg/WatchYourLAN/internal/models"
	"github.com/aceberg/WatchYourLAN/internal/scan"
)

func configHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig
	guiData.Icon = Icon

	guiData.Themes = []string{"cerulean", "cosmo", "cyborg", "darkly", "flatly", "journal", "litera", "lumen", "lux", "materia", "minty", "morph", "pulse", "quartz", "sandstone", "simplex", "sketchy", "slate", "solar", "spacelab", "superhero", "united", "vapor", "yeti", "zephyr"}

	tmpl, err := template.ParseFS(TemplHTML, TemplPath+"config.html", TemplPath+"header.html", TemplPath+"footer.html")
	check.IfError(err)
	err = tmpl.ExecuteTemplate(w, "header", guiData)
	check.IfError(err)
	err = tmpl.ExecuteTemplate(w, "config", guiData)
	check.IfError(err)
}

func saveConfigHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	AppConfig.Iface = r.FormValue("iface")
	AppConfig.GuiIP = r.FormValue("host")
	AppConfig.GuiPort = r.FormValue("port")
	AppConfig.ShoutURL = r.FormValue("shout")
	AppConfig.Theme = r.FormValue("theme")

	timeout := r.FormValue("timeout")
	AppConfig.Timeout, err = strconv.Atoi(timeout)
	check.IfError(err)

	close(QuitScan)

	conf.Write(ConfigPath, AppConfig)

	QuitScan = make(chan bool)
	go scan.Start(AppConfig, QuitScan)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}
