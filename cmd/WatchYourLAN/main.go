package main

import (
	"flag"

	"github.com/aceberg/WatchYourLAN/internal/check"
	"github.com/aceberg/WatchYourLAN/internal/web"
)

const configPath = "/data/config"
const bootPath = "https://cdn.jsdelivr.net/npm/bootswatch@5.2.3/dist"
// const bootPath = "file://data/node_modules/bootswatch/dist"

func main() {
	confPtr := flag.String("c", configPath, "Path to config file")
	bootPtr := flag.String("b", bootPath, "Path to Bootswatch")
	flag.Parse()

	check.Path(*confPtr)

	web.Gui(*confPtr, *bootPtr) // Start web GUI
}
