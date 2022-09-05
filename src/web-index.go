package main

import (
  "fmt"
  "log"
  "net/http"
  "html/template"
  "strconv"
)

func index(w http.ResponseWriter, r *http.Request) {
	type allData struct {
		Config Conf
		Hosts []Host
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

func webgui() {
	// fmt.Println(FoundHosts)
	address := AppConfig.GuiIP + ":" + AppConfig.GuiPort

	log.Println("=================================== ")
	log.Println(fmt.Sprintf("Web GUI at http://%s", address))
	log.Println("=================================== ")

	http.HandleFunc("/", index)
	http.HandleFunc("/home/", home)
	http.HandleFunc("/offline/", offline)
	http.HandleFunc("/online/", online)
	http.HandleFunc("/search_hosts/", search_hosts)
	http.HandleFunc("/sort_hosts/", sort_hosts)
	http.HandleFunc("/theme/", theme)
	http.HandleFunc("/update_host/", update_host)
	http.ListenAndServe(address, nil)
}