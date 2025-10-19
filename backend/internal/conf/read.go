package conf

import (
	"strings"

	"github.com/spf13/viper"

	"github.com/aceberg/WatchYourLAN/internal/check"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

func read(path string) (config models.Conf) {

	viper.SetDefault("HOST", "0.0.0.0")
	viper.SetDefault("PORT", "8840")
	viper.SetDefault("THEME", "sand")
	viper.SetDefault("COLOR", "dark")
	viper.SetDefault("NODEPATH", "")
	viper.SetDefault("LOG_LEVEL", "info")
	viper.SetDefault("ARP_ARGS", "")
	viper.SetDefault("ARP_STRS_JOINED", "")
	viper.SetDefault("IFACES", "")
	viper.SetDefault("TIMEOUT", 120)
	viper.SetDefault("TRIM_HIST", 48)
	viper.SetDefault("SHOUTRRR_URL", "")

	viper.SetDefault("USE_DB", "sqlite")
	viper.SetDefault("PG_CONNECT", "")

	viper.SetDefault("INFLUX_ENABLE", false)

	viper.SetDefault("PROMETHEUS_ENABLE", false)

	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	check.IfError(err)

	viper.AutomaticEnv() // Get ENVIRONMENT variables

	config.Host = stripQuotes(viper.GetString("HOST"))
	config.Port = stripQuotes(viper.GetString("PORT"))
	config.Theme = stripQuotes(viper.GetString("THEME"))
	config.Color = stripQuotes(viper.GetString("COLOR"))
	config.NodePath = stripQuotes(viper.GetString("NODEPATH"))
	config.LogLevel = stripQuotes(viper.GetString("LOG_LEVEL"))
	config.ArpArgs = stripQuotes(viper.GetString("ARP_ARGS"))
	config.ArpStrs = viper.GetStringSlice("ARP_STRS")
	config.Ifaces = stripQuotes(viper.GetString("IFACES"))
	config.Timeout = viper.GetInt("TIMEOUT")
	config.TrimHist = viper.GetInt("TRIM_HIST")
	config.ShoutURL = stripQuotes(viper.GetString("SHOUTRRR_URL"))

	config.UseDB = stripQuotes(viper.GetString("USE_DB"))
	config.PGConnect = stripQuotes(viper.GetString("PG_CONNECT"))

	config.InfluxEnable = viper.GetBool("INFLUX_ENABLE")
	config.InfluxSkipTLS = viper.GetBool("INFLUX_SKIP_TLS")
	config.InfluxAddr = stripQuotes(viper.GetString("INFLUX_ADDR"))
	config.InfluxToken = stripQuotes(viper.GetString("INFLUX_TOKEN"))
	config.InfluxOrg = stripQuotes(viper.GetString("INFLUX_ORG"))
	config.InfluxBucket = stripQuotes(viper.GetString("INFLUX_BUCKET"))

	config.PrometheusEnable = viper.GetBool("PROMETHEUS_ENABLE")

	joined := stripQuotes(viper.GetString("ARP_STRS_JOINED"))
	// slog.Info("ARP_STRS_JOINED: " + joined)

	if joined != "" {
		splitStrs := strings.Split(joined, ",")
		config.ArpStrs = make([]string, len(splitStrs))
		for i, str := range splitStrs {
			config.ArpStrs[i] = stripQuotes(strings.TrimSpace(str))
		}
	}

	// Also clean quotes from ArpStrs if they were set directly
	for i, str := range config.ArpStrs {
		config.ArpStrs[i] = stripQuotes(str)
	}

	return config
}
