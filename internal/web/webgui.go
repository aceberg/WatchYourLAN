package web

import (
	"log"
	"net/http"

	"github.com/aceberg/WatchYourLAN/internal/check"
	"github.com/aceberg/WatchYourLAN/internal/conf"
	"github.com/aceberg/WatchYourLAN/internal/db"
	"github.com/aceberg/WatchYourLAN/internal/scan"
)

// Gui - start web GUI
func Gui(configPath, nodePath string) {

	ConfigPath = configPath
	AppConfig = conf.Get(ConfigPath)
	AppConfig.NodePath = nodePath
	AppConfig.Icon = Icon

	address := AppConfig.GuiIP + ":" + AppConfig.GuiPort

	db.Create(AppConfig.DbPath)

	QuitScan = make(chan bool)

	go scan.Start(AppConfig, QuitScan)

	AllHosts = db.Select(AppConfig.DbPath)

	log.Println("=================================== ")
	log.Printf("Web GUI at http://%s", address)
	log.Println("=================================== ")

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/clear/", clearHandler)
	http.HandleFunc("/config/", configHandler)
	http.HandleFunc("/del_host/", delHandler)
	http.HandleFunc("/home/", homeHandler)
	http.HandleFunc("/host/", hostHandler)
	http.HandleFunc("/line/", lineHandler)
	http.HandleFunc("/port_scan/", portHandler)
	http.HandleFunc("/save_config/", saveConfigHandler)
	http.HandleFunc("/search_hosts/", searchHandler)
	http.HandleFunc("/sort_hosts/", sortHandler)
	http.HandleFunc("/test_notify/", testNotifyHandler)
	http.HandleFunc("/update_host/", updateHandler)
	err := http.ListenAndServe(address, nil)
	check.IfError(err)
}
