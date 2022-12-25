package web

import (
	// "embed"
	"log"
	"net/http"

	"github.com/aceberg/WatchYourLAN/internal/models"
)

var (
	// AppConfig - app config
	AppConfig models.Conf

	// TemplPath - path to html templates
	TemplPath string
)

// var AllHosts []Host

////go:embed templates/*
// var TemplHTML embed.FS

// Gui - start web GUI
func Gui(appConfig models.Conf) {

	TemplPath = "../../internal/web/templates/"

	AppConfig = appConfig
	address := AppConfig.GuiIP + ":" + AppConfig.GuiPort

	log.Println("=================================== ")
	log.Printf("Web GUI at http://%s", address)
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
