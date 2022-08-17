package main

import (
  "fmt"
  "net/http"
  "html/template"
)

var FoundHosts []Host

func index (w http.ResponseWriter, r *http.Request) {
	fmt.Println(FoundHosts)
	tmpl, _ := template.ParseFiles("templates/index.html", "templates/header.html")
	tmpl.ExecuteTemplate(w, "index", FoundHosts)
}

func webgui (config Conf, hosts []Host) {
	// fmt.Println(FoundHosts)
	FoundHosts = hosts
	address := config.GuiIP + ":" + config.GuiPort

	http.HandleFunc("/", index)
	http.ListenAndServe(address, nil)
}