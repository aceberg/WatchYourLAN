package main

import (
	"fmt"

	"github.com/aceberg/WatchYourLAN/internal/conf"
	"github.com/aceberg/WatchYourLAN/internal/db"
	"github.com/aceberg/WatchYourLAN/internal/web"
)

const configPath = "/data/config"

func main() {
	appConfig := conf.Get(configPath) // Get config from Defaults, Config file, Env

	fmt.Println("CONF =", appConfig)

	db.Create(appConfig.DbPath) // Check if DB exists. Create if not

	web.Gui(appConfig) // Start web GUI
}
