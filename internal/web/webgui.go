package web

import (
	// "embed"
	"fmt"
	"log"
	"net/http"

	"github.com/aceberg/WatchYourLAN/internal/models"
)

var (
	AppConfig models.Conf

	TemplPath string
)

// var AllHosts []Host

////go:embed templates/*
// var TemplHTML embed.FS

func Webgui(appConfig models.Conf) {

	TemplPath = "../../internal/web/templates/"

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
