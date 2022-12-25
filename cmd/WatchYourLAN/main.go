package main

import (
	"fmt"

	"github.com/aceberg/WatchYourLAN/internal/conf"
	"github.com/aceberg/WatchYourLAN/internal/db"
	"github.com/aceberg/WatchYourLAN/internal/web"
)

const configPath = "/data/config"

func main() {
	appConfig := conf.GetConfig(configPath) // Get config from Defaults, Config file, Env

	fmt.Println("CONF =", appConfig)

	db.CreateDB(appConfig.DbPath) // Check if DB exists. Create if not

	web.Webgui(appConfig) // Start web GUI
}
