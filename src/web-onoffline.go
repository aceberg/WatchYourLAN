package main

import (
  "net/http"
  "html/template"
)

func offline(w http.ResponseWriter, r *http.Request) {
	type allData struct {
		Config Conf
		Hosts []Host
	}
	var guiData allData
	guiData.Config = AppConfig
	guiData.Hosts = []Host{}
	
	for _, oneHost := range AllHosts {
		if oneHost.Now == 0 {
			guiData.Hosts = append(guiData.Hosts, oneHost)
		}
	}

	tmpl, _ := template.ParseFiles("templates/offline.html", "templates/header.html", "templates/footer.html")
	tmpl.ExecuteTemplate(w, "offline", guiData)
}

func online(w http.ResponseWriter, r *http.Request) {
	type allData struct {
		Config Conf
		Hosts []Host
	}
	var guiData allData
	guiData.Config = AppConfig
	guiData.Hosts = []Host{}
	
	for _, oneHost := range AllHosts {
		if oneHost.Now == 1 {
			guiData.Hosts = append(guiData.Hosts, oneHost)
		}
	}

	tmpl, _ := template.ParseFiles("templates/offline.html", "templates/header.html", "templates/footer.html")
	tmpl.ExecuteTemplate(w, "offline", guiData)
}