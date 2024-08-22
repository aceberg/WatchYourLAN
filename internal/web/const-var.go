package web

import (
	"embed"

	"github.com/aceberg/WatchYourLAN/internal/models"
)

var (
	// appConfig - config for Web Gui
	appConfig models.Conf

	allHosts  []models.Host
	histHosts []models.Host

	quitScan chan bool
)

// templFS - html templates
//
//go:embed templates/*
var templFS embed.FS

// pubFS - public folder
//
//go:embed public/*
var pubFS embed.FS
