package models

// Conf - app config
type Conf struct {
	Host     string
	Port     string
	Theme    string
	Color    string
	DirPath  string
	ConfPath string
	DBPath   string
	NodePath string
	Ifaces   string
	// Scaner   string
	ArpArgs  string
	Timeout  int
	IgnoreIP string
}

// Host - one host
type Host struct {
	ID    int    `db:"ID"`
	Name  string `db:"NAME"`
	DNS   string `db:"DNS"`
	Iface string `db:"IFACE"`
	IP    string `db:"IP"`
	Mac   string `db:"MAC"`
	Hw    string `db:"HW"`
	Date  string `db:"DATE"`
	Known int    `db:"KNOWN"`
	Now   int    `db:"NOW"`
}

// GuiData - all data sent to html page
type GuiData struct {
	Config  Conf
	Host    Host
	Themes  []string
	Version string
}
