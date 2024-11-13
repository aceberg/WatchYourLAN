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
	LogLevel string
	Ifaces   string
	ArpArgs  string
	ArpStrs  []string
	Timeout  int
	TrimHist int
	HistInDB bool
	ShoutURL string
	// PostgreSQL
	UseDB     string
	PGConnect string
	// InfluxDB
	InfluxEnable  bool
	InfluxAddr    string
	InfluxToken   string
	InfluxOrg     string
	InfluxBucket  string
	InfluxSkipTLS bool
}

// Host - one host
type Host struct {
	ID        int    `db:"ID"`
	Name      string `db:"NAME"`
	DNS       string `db:"DNS"`
	Iface     string `db:"IFACE"`
	IP        string `db:"IP"`
	Mac       string `db:"MAC"`
	Hw        string `db:"HW"`
	Date      string `db:"DATE"`
	FirstSeen string `db:"FIRSTSEEN"`
	Known     int    `db:"KNOWN"`
	Now       int    `db:"NOW"`
}

// Stat - status
type Stat struct {
	Total   int
	Online  int
	Offline int
	Known   int
	Unknown int
}

// GuiData - all data sent to html page
type GuiData struct {
	Config  Conf
	Host    Host
	Themes  []string
	Version string
}
