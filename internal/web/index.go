package web

import (
	//   "fmt"
	"html/template"
	"net/http"

	"github.com/aceberg/WatchYourLAN/internal/models"
)

func index(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData
	guiData.Config = AppConfig
	guiData.Hosts = []models.Host{}

	tmpl, _ := template.ParseFiles(TemplPath+"index.html", TemplPath+"header.html", TemplPath+"footer.html")
	tmpl.ExecuteTemplate(w, "index", guiData)
}
