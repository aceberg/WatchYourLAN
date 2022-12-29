package web

import (
	"html/template"
	"net/http"

	"github.com/aceberg/WatchYourLAN/internal/check"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

func offlineHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig
	guiData.Hosts = []models.Host{}
	guiData.Icon = Icon

	for _, oneHost := range AllHosts {
		if oneHost.Now == 0 {
			guiData.Hosts = append(guiData.Hosts, oneHost)
		}
	}

	tmpl, _ := template.ParseFS(TemplHTML, TemplPath+"offline.html", TemplPath+"header.html", TemplPath+"footer.html")
	err := tmpl.ExecuteTemplate(w, "header", guiData)
	check.IfError(err)
	err = tmpl.ExecuteTemplate(w, "offline", guiData)
	check.IfError(err)
}

func onlineHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig
	guiData.Hosts = []models.Host{}
	guiData.Icon = Icon

	for _, oneHost := range AllHosts {
		if oneHost.Now == 1 {
			guiData.Hosts = append(guiData.Hosts, oneHost)
		}
	}

	tmpl, _ := template.ParseFS(TemplHTML, TemplPath+"offline.html", TemplPath+"header.html", TemplPath+"footer.html")
	err := tmpl.ExecuteTemplate(w, "header", guiData)
	check.IfError(err)
	err = tmpl.ExecuteTemplate(w, "offline", guiData)
	check.IfError(err)
}
