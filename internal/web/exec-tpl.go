package web

import (
	"html/template"
	"net/http"

	"github.com/aceberg/WatchYourLAN/internal/check"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

func execTemplate(w http.ResponseWriter, tpl string, guiData models.GuiData) {
	// tmpl, err := template.ParseFiles(TemplPath+tpl+".html", TemplPath+"header.html", TemplPath+"footer.html")
	tmpl, err := template.ParseFS(TemplHTML, TemplPath+tpl+".html", TemplPath+"header.html", TemplPath+"footer.html")
	check.IfError(err)
	err = tmpl.ExecuteTemplate(w, "header", guiData)
	check.IfError(err)
	err = tmpl.ExecuteTemplate(w, tpl, guiData)
	check.IfError(err)
}
