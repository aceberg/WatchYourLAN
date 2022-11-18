package web

import (
	//   "fmt"
	"html/template"
	"net/http"
	//   "strconv"
	. "github.com/aceberg/WatchYourLAN/pkg/models"
)

func index(w http.ResponseWriter, r *http.Request) {
	var guiData GuiData
	guiData.Config = AppConfig
	guiData.Hosts = []Host{}

	tmpl, _ := template.ParseFS(TemplHTML, "templates/index.html", "templates/header.html", "templates/footer.html")
	tmpl.ExecuteTemplate(w, "index", guiData)
}

// func update_host(w http.ResponseWriter, r *http.Request) {
// 	idStr := r.FormValue("id")
// 	name := r.FormValue("name")
// 	knownStr := r.FormValue("known")

// 	if idStr == "" {
// 		fmt.Fprintf(w, "No data!")
// 	} else {
// 		var known uint16
// 		id, _ := strconv.Atoi(idStr)
// 		known = 0
// 		if knownStr == "on" {
// 			known = 1
// 		}

// 		for i, oneHost := range AllHosts {
// 			if oneHost.Id == uint16(id) {
// 				AllHosts[i].Name = name
// 				AllHosts[i].Known = known
// 				db_update(AllHosts[i])
// 			}
// 		}
// 	}

// 	http.Redirect(w, r, r.Header.Get("Referer"), 302)
// }
