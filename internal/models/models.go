package models

// Conf - app config
type Conf struct {
	Host     string
	Port     string
	Theme    string
	Color    string
	DirPath  string
	ConfPath string
	NodePath string
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
	Hist    []History
}
