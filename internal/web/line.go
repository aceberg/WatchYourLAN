package web

import (
	"html/template"
	"net/http"

	"github.com/aceberg/WatchYourLAN/internal/check"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

func lineHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig
	guiData.Hosts = []models.Host{}
	guiData.Icon = Icon

	state := r.URL.Query().Get("state")

	for _, oneHost := range AllHosts {
		if state == "off" && oneHost.Now == 0 {
			guiData.Hosts = append(guiData.Hosts, oneHost)
		}
		if state == "on" && oneHost.Now == 1 {
			guiData.Hosts = append(guiData.Hosts, oneHost)
		}
	}

	tmpl, _ := template.ParseFS(TemplHTML, TemplPath+"line.html", TemplPath+"header.html", TemplPath+"footer.html")
	err := tmpl.ExecuteTemplate(w, "header", guiData)
	check.IfError(err)
	err = tmpl.ExecuteTemplate(w, "line", guiData)
	check.IfError(err)
}
