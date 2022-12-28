package web

import (
	"embed"
	"log"
	"net/http"

	"github.com/aceberg/WatchYourLAN/internal/check"
	"github.com/aceberg/WatchYourLAN/internal/db"
	"github.com/aceberg/WatchYourLAN/internal/models"
	"github.com/aceberg/WatchYourLAN/internal/scan"
)

var (
	// AppConfig - app config
	AppConfig models.Conf

	// AllHosts - all hosts from DB
	AllHosts []models.Host
)

// TemplHTML - embed templates
//
//go:embed templates/*
var TemplHTML embed.FS

// TemplPath - path to html templates
const TemplPath = "templates/"

// const TemplPath = "../../internal/web/templates/"

// Gui - start web GUI
func Gui(appConfig models.Conf) {
	AppConfig = appConfig
	address := AppConfig.GuiIP + ":" + AppConfig.GuiPort

	go scan.Start(AppConfig)

	log.Println("=================================== ")
	log.Printf("Web GUI at http://%s", address)
	log.Println("=================================== ")

	AllHosts = db.Select(AppConfig.DbPath)

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/home/", homeHandler)
	http.HandleFunc("/offline/", offlineHandler)
	http.HandleFunc("/online/", onlineHandler)
	http.HandleFunc("/search_hosts/", searchHandler)
	http.HandleFunc("/sort_hosts/", sortHandler)
	// http.HandleFunc("/theme/", theme)
	http.HandleFunc("/update_host/", updateHandler)
	err := http.ListenAndServe(address, nil)
	check.IfError(err)
}
