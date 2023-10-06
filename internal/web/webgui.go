package web

import (
	"log"
	"net/http"

	"github.com/aceberg/WatchYourLAN/internal/auth"
	"github.com/aceberg/WatchYourLAN/internal/check"
	"github.com/aceberg/WatchYourLAN/internal/conf"
	"github.com/aceberg/WatchYourLAN/internal/db"
	"github.com/aceberg/WatchYourLAN/internal/migrate"
	"github.com/aceberg/WatchYourLAN/internal/scan"
)

// Gui - start web GUI
func Gui(configPath, nodePath string) {

	ConfigPath = migrate.ToYAML(configPath)
	AppConfig, authConf = conf.Get(ConfigPath)
	AppConfig.NodePath = nodePath
	AppConfig.Icon = Icon

	address := AppConfig.GuiIP + ":" + AppConfig.GuiPort

	QuitScan = make(chan bool)
	go scan.Start(AppConfig, QuitScan)
	go trimHistoryRoutine() // trim-history.go

	updateAllHosts() // webgui.go

	log.Println("=================================== ")
	log.Printf("Web GUI at http://%s", address)
	log.Println("=================================== ")

	http.HandleFunc("/login/", loginHandler) // login.go

	http.HandleFunc("/", auth.Auth(indexHandler, &authConf))
	http.HandleFunc("/auth_conf/", auth.Auth(authConfHandler, &authConf))     // auth.go
	http.HandleFunc("/auth_save/", auth.Auth(saveAuthHandler, &authConf))     // auth.go
	http.HandleFunc("/clear/", auth.Auth(clearHandler, &authConf))            // config.go
	http.HandleFunc("/config/", auth.Auth(configHandler, &authConf))          // config.go
	http.HandleFunc("/del_host/", auth.Auth(delHandler, &authConf))           // host.go
	http.HandleFunc("/history/", auth.Auth(historyHandler, &authConf))        // history.go
	http.HandleFunc("/host/", auth.Auth(hostHandler, &authConf))              // host.go
	http.HandleFunc("/line/", auth.Auth(lineHandler, &authConf))              // line.go
	http.HandleFunc("/port_scan/", auth.Auth(portHandler, &authConf))         // port.go
	http.HandleFunc("/save_config/", auth.Auth(saveConfigHandler, &authConf)) // config.go
	http.HandleFunc("/search_hosts/", auth.Auth(searchHandler, &authConf))    // search.go
	http.HandleFunc("/sort_hosts/", auth.Auth(sortHandler, &authConf))        // sort.go
	http.HandleFunc("/test_notify/", auth.Auth(testNotifyHandler, &authConf)) // config.go
	http.HandleFunc("/update_host/", auth.Auth(updateHandler, &authConf))     // update.go
	err := http.ListenAndServe(address, nil)
	check.IfError(err)
}

func updateAllHosts() {
	AllHosts = db.Select(AppConfig.DbPath)
}
