package models

import (
	"github.com/aceberg/WatchYourLAN/internal/auth"
)

// Host - one host
type Host struct {
	ID    uint16
	Name  string
	IP    string
	Mac   string
	Hw    string
	Date  string
	Known uint16
	Now   uint16
}

// Conf - app config
type Conf struct {
	Iface    string
	DbPath   string
	GuiIP    string
	GuiPort  string
	Timeout  int
	ShoutURL string
	Theme    string
	Color    string
	IgnoreIP string
	LogLevel string
	NodePath string
	Icon     string
	Auth     bool
}

// GuiData - all data sent to html page
type GuiData struct {
	Config  Conf
	Hosts   []Host
	Themes  []string
	Version string
	Auth    auth.Conf
}
