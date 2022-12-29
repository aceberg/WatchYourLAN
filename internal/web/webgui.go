package web

import (
	"embed"
	"log"
	"net/http"

	"github.com/aceberg/WatchYourLAN/internal/check"
	"github.com/aceberg/WatchYourLAN/internal/conf"
	"github.com/aceberg/WatchYourLAN/internal/db"
	"github.com/aceberg/WatchYourLAN/internal/models"
	"github.com/aceberg/WatchYourLAN/internal/scan"
)

var (
	// AppConfig - app config
	AppConfig models.Conf

	// AllHosts - all hosts from DB
	AllHosts []models.Host

	// ConfigPath - path to config file
	ConfigPath string

	// QuitScan - send stop signal to scan
	QuitScan chan bool

	// TemplHTML - embed templates
	//
	//go:embed templates/*
	TemplHTML embed.FS
)

// TemplPath - path to html templates
const TemplPath = "templates/"

// Gui - start web GUI
func Gui(configPath string) {

	ConfigPath = configPath
	AppConfig = conf.Get(ConfigPath)

	address := AppConfig.GuiIP + ":" + AppConfig.GuiPort

	db.Create(AppConfig.DbPath)

	QuitScan = make(chan bool)

	go scan.Start(AppConfig, QuitScan)

	AllHosts = db.Select(AppConfig.DbPath)

	log.Println("=================================== ")
	log.Printf("Web GUI at http://%s", address)
	log.Println("=================================== ")

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/config/", configHandler)
	http.HandleFunc("/home/", homeHandler)
	http.HandleFunc("/offline/", offlineHandler)
	http.HandleFunc("/online/", onlineHandler)
	http.HandleFunc("/save_config/", saveConfigHandler)
	http.HandleFunc("/search_hosts/", searchHandler)
	http.HandleFunc("/sort_hosts/", sortHandler)
	http.HandleFunc("/update_host/", updateHandler)
	err := http.ListenAndServe(address, nil)
	check.IfError(err)
}
