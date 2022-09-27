package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func index(w http.ResponseWriter, r *http.Request) {
	type allData struct {
		Config Conf
		Hosts  []Host
	}
	var guiData allData
	guiData.Config = AppConfig
	guiData.Hosts = AllHosts

	tmpl, _ := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	tmpl.ExecuteTemplate(w, "index", guiData)
}

func update_host(w http.ResponseWriter, r *http.Request) {
	idStr := r.FormValue("id")
	name := r.FormValue("name")
	knownStr := r.FormValue("known")

	if idStr == "" {
		fmt.Fprintf(w, "No data!")
	} else {
		var known uint16
		id, _ := strconv.Atoi(idStr)
		known = 0
		if knownStr == "on" {
			known = 1
		}

		for i, oneHost := range AllHosts {
			if oneHost.Id == uint16(id) {
				AllHosts[i].Name = name
				AllHosts[i].Known = known
				db_update(AllHosts[i])
			}
		}
	}

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}

func basicAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if AppConfig.GuiAuth == "" {
			next.ServeHTTP(w, r)
			return
		}

		username, password, ok := r.BasicAuth()
		if ok {
			userCredentials := fmt.Sprintf(`%s:%s`, username, password)
			if userCredentials == AppConfig.GuiAuth {
				next.ServeHTTP(w, r)
				return
			}
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}

func webgui() {
	// fmt.Println(FoundHosts)
	address := AppConfig.GuiIP + ":" + AppConfig.GuiPort

	log.Println("=================================== ")
	log.Println(fmt.Sprintf("Web GUI at http://%s", address))
	log.Println("=================================== ")

	http.HandleFunc("/", basicAuth(index))
	http.HandleFunc("/home/", basicAuth(home))
	http.HandleFunc("/offline/", basicAuth(offline))
	http.HandleFunc("/online/", basicAuth(online))
	http.HandleFunc("/search_hosts/", basicAuth(search_hosts))
	http.HandleFunc("/sort_hosts/", basicAuth(sort_hosts))
	http.HandleFunc("/theme/", basicAuth(theme))
	http.HandleFunc("/update_host/", basicAuth(update_host))
	http.ListenAndServe(address, nil)
}
