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
}

// Host - one host
type Host struct {
	ID    uint16 `db:"ID"`
	Name  string `db:"NAME"`
	IP    string `db:"IP"`
	Mac   string `db:"MAC"`
	Hw    string `db:"HW"`
	Date  string `db:"DATE"`
	Known uint16 `db:"KNOWN"`
	Now   uint16 `db:"NOW"`
}

// // History for hosts
// type History struct {
// 	ID 		int
// 	HostID	int
// 	Name	string
// 	IP	string
// 	Date string
// 	State bool
// }

// GuiData - all data sent to html page
type GuiData struct {
	Config  Conf
	Hosts   []Host
	Themes  []string
	Version string
	Auth    auth.Conf
	// HistLog []History
}
