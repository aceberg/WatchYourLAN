package web

import (
	//   "fmt"
	"html/template"
	"net/http"

	"github.com/aceberg/WatchYourLAN/internal/check"
	"github.com/aceberg/WatchYourLAN/internal/db"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig
	guiData.Hosts = AllHosts
	guiData.Icon = Icon

	tmpl, _ := template.ParseFiles(TemplPath+"index.html", TemplPath+"header.html", TemplPath+"footer.html")
	err := tmpl.ExecuteTemplate(w, "header", guiData)
	check.IfError(err)
	err = tmpl.ExecuteTemplate(w, "index", guiData)
	check.IfError(err)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	AllHosts = db.Select(AppConfig.DbPath)
	http.Redirect(w, r, "/", 302)
}
