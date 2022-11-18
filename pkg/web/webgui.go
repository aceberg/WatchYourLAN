package web

import (
	"embed"
	"fmt"
	. "github.com/aceberg/WatchYourLAN/pkg/models"
	"log"
	"net/http"
)

var AppConfig Conf

// var AllHosts []Host

//go:embed templates/*
var TemplHTML embed.FS

func Webgui(appConfig Conf) {

	AppConfig = appConfig
	address := AppConfig.GuiIP + ":" + AppConfig.GuiPort

	log.Println("=================================== ")
	log.Println(fmt.Sprintf("Web GUI at http://%s", address))
	log.Println("=================================== ")

	http.HandleFunc("/", index)
	http.HandleFunc("/home/", index)
	// http.HandleFunc("/offline/", offline)
	// http.HandleFunc("/online/", online)
	// http.HandleFunc("/search_hosts/", search_hosts)
	// http.HandleFunc("/sort_hosts/", sort_hosts)
	// http.HandleFunc("/theme/", theme)
	// http.HandleFunc("/update_host/", update_host)
	http.ListenAndServe(address, nil)
}
