package models

import (
	"github.com/aceberg/WatchYourLAN/internal/auth"
)

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
	HistDays string
}

// Host - one host
type Host struct {
	ID    int    `db:"ID"`
	Name  string `db:"NAME"`
	IP    string `db:"IP"`
	Mac   string `db:"MAC"`
	Hw    string `db:"HW"`
	Date  string `db:"DATE"`
	Known int    `db:"KNOWN"`
	Now   int    `db:"NOW"`
}

// History for hosts
type History struct {
	ID    int    `db:"ID"`
	Host  int    `db:"HOST"`
	Name  string `db:"NAME"`
	IP    string `db:"IP"`
	Mac   string `db:"MAC"`
	Hw    string `db:"HW"`
	Date  string `db:"DATE"`
	Known int    `db:"KNOWN"`
	State int    `db:"STATE"`
}

// GuiData - all data sent to html page
type GuiData struct {
	Config  Conf
	Hosts   []Host
	Themes  []string
	Version string
	Auth    auth.Conf
	Hist    []History
}
