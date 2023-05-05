package models

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
	IgnoreIP string
	LogLevel string
	BootPath string
	Icon     string
}

// GuiData - all data sent to html page
type GuiData struct {
	Config  Conf
	Hosts   []Host
	Themes  []string
	Version string
	Release string
}
