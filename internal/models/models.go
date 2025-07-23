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
	// Prometheus
	PrometheusEnable bool
}

// Host - one host
type Host struct {
	ID    int    `gorm:"column:ID;primaryKey"`
	Name  string `gorm:"column:NAME"`
	DNS   string `gorm:"column:DNS"`
	Iface string `gorm:"column:IFACE"`
	IP    string `gorm:"column:IP"`
	Mac   string `gorm:"column:MAC"`
	Hw    string `gorm:"column:HW"`
	Date  string `gorm:"column:DATE"`
	Known int    `gorm:"column:KNOWN"`
	Now   int    `gorm:"column:NOW"`
}

// Stat - status
type Stat struct {
	Total   int
	Online  int
	Offline int
	Known   int
	Unknown int
}
