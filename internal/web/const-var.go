package web

import (
	"embed"

	"github.com/aceberg/WatchYourLAN/internal/models"
)

var (
	// appConfig - config for Web Gui
	appConfig models.Conf
)

// templFS - html templates
//
//go:embed templates/*
var templFS embed.FS

// pubFS - public folder
//
//go:embed public/*
var pubFS embed.FS
