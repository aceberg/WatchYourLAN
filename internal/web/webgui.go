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

	AllHosts = db.Select(AppConfig.DbPath)

	log.Println("=================================== ")
	log.Printf("Web GUI at http://%s", address)
	log.Println("=================================== ")

	http.HandleFunc("/login/", loginHandler) // login.go

	http.HandleFunc("/", auth.Auth(indexHandler, &authConf))
	http.HandleFunc("/auth_conf/", auth.Auth(authConfHandler, &authConf)) // auth-conf.go
	http.HandleFunc("/auth_save/", auth.Auth(saveAuthHandler, &authConf)) // auth-conf.go
	http.HandleFunc("/clear/", auth.Auth(clearHandler, &authConf))        // config.go
	http.HandleFunc("/config/", auth.Auth(configHandler, &authConf))      // config.go
	http.HandleFunc("/del_host/", auth.Auth(delHandler, &authConf))
	http.HandleFunc("/host/", auth.Auth(hostHandler, &authConf))
	http.HandleFunc("/line/", auth.Auth(lineHandler, &authConf))
	http.HandleFunc("/port_scan/", auth.Auth(portHandler, &authConf))
	http.HandleFunc("/save_config/", auth.Auth(saveConfigHandler, &authConf))
	http.HandleFunc("/search_hosts/", auth.Auth(searchHandler, &authConf))
	http.HandleFunc("/sort_hosts/", auth.Auth(sortHandler, &authConf))
	http.HandleFunc("/test_notify/", auth.Auth(testNotifyHandler, &authConf))
	http.HandleFunc("/update_host/", auth.Auth(updateHandler, &authConf))
	err := http.ListenAndServe(address, nil)
	check.IfError(err)
}
