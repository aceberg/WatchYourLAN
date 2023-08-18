package web

import (
	"embed"

	"github.com/aceberg/WatchYourLAN/internal/auth"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

var (
	authConf auth.Conf

	// AppConfig - app config
	AppConfig models.Conf

	// AllHosts - all hosts from DB
	AllHosts []models.Host

	// ConfigPath - path to config file
	ConfigPath string

	// QuitScan - send stop signal to scan
	QuitScan chan bool

	// TemplHTML - embed templates
	//
	//go:embed templates/*
	TemplHTML embed.FS
)

// TemplPath - path to html templates
const TemplPath = "templates/"
