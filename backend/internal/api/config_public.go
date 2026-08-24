package api

import "github.com/aceberg/WatchYourLAN/internal/models"

type publicConfig struct {
	Host                  string
	Port                  string
	Theme                 string
	Color                 string
	DirPath               string
	NodePath              string
	LogLevel              string
	Ifaces                string
	ArpArgs               string
	ArpStrs               []string
	Timeout               int
	TrimHist              int
	ConnectivityRetention int
	ShoutURL              string
	ShoutURLConfigured    bool
	UseDB                 string
	PGConnect             string
	PGConnectConfigured   bool
	InfluxEnable          bool
	InfluxAddr            string
	InfluxToken           string
	InfluxTokenConfigured bool
	InfluxOrg             string
	InfluxBucket          string
	InfluxSkipTLS         bool
	PrometheusEnable      bool
}

func toPublicConfig(config models.Conf) publicConfig {
	return publicConfig{
		Host:                  config.Host,
		Port:                  config.Port,
		Theme:                 config.Theme,
		Color:                 config.Color,
		DirPath:               config.DirPath,
		NodePath:              config.NodePath,
		LogLevel:              config.LogLevel,
		Ifaces:                config.Ifaces,
		ArpArgs:               config.ArpArgs,
		ArpStrs:               append([]string(nil), config.ArpStrs...),
		Timeout:               config.Timeout,
		TrimHist:              config.TrimHist,
		ConnectivityRetention: config.ConnectivityRetention,
		ShoutURL:              "",
		ShoutURLConfigured:    config.ShoutURL != "",
		UseDB:                 config.UseDB,
		PGConnect:             "",
		PGConnectConfigured:   config.PGConnect != "",
		InfluxEnable:          config.InfluxEnable,
		InfluxAddr:            config.InfluxAddr,
		InfluxToken:           "",
		InfluxTokenConfigured: config.InfluxToken != "",
		InfluxOrg:             config.InfluxOrg,
		InfluxBucket:          config.InfluxBucket,
		InfluxSkipTLS:         config.InfluxSkipTLS,
		PrometheusEnable:      config.PrometheusEnable,
	}
}
