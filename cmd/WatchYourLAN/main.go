package main

import (
	"flag"

	"github.com/aceberg/WatchYourLAN/internal/check"
	"github.com/aceberg/WatchYourLAN/internal/web"
)

const configPath = "/data/config"

func main() {
	confPtr := flag.String("c", configPath, "Path to config file")
	flag.Parse()

	check.Path(*confPtr)

	web.Gui(*confPtr) // Start web GUI
}
