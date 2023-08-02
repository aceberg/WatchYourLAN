package main

import (
	"flag"

	"github.com/aceberg/WatchYourLAN/internal/check"
	"github.com/aceberg/WatchYourLAN/internal/web"
)

const configPath = "/data/config.yaml"
const nodePath = ""

func main() {
	confPtr := flag.String("c", configPath, "Path to config file")
	nodePtr := flag.String("n", nodePath, "Path to node modules")
	flag.Parse()

	check.Path(*confPtr)

	web.Gui(*confPtr, *nodePtr) // Start web GUI
}
