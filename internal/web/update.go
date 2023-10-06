package web

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/aceberg/WatchYourLAN/internal/db"
)

func updateHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.FormValue("id")
	name := r.FormValue("name")
	knownStr := r.FormValue("known")

	if idStr == "" {
		fmt.Fprintf(w, "No data!")
	} else {
		var known int
		id, _ := strconv.Atoi(idStr)
		known = 0
		if knownStr == "on" {
			known = 1
		}

		for i, oneHost := range AllHosts {
			if oneHost.ID == id {
				AllHosts[i].Name = name
				AllHosts[i].Known = known
				db.Update(AppConfig.DbPath, AllHosts[i])
			}
		}
	}

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}
