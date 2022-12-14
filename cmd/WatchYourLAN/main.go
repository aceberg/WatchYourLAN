package main

import (
	"flag"

	"github.com/aceberg/WatchYourLAN/internal/check"
	"github.com/aceberg/WatchYourLAN/internal/web"
)

const configPath = "/data/config"

const bootPath = ""

func main() {
	confPtr := flag.String("c", configPath, "Path to config file")
	bootPtr := flag.String("b", bootPath, "Path to local Bootswatch")
	flag.Parse()

	check.Path(*confPtr)

	web.Gui(*confPtr, *bootPtr) // Start web GUI
}
