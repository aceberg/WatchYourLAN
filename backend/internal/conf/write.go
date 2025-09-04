package conf

import (
	"log/slog"

	"github.com/spf13/viper"

	"github.com/aceberg/WatchYourLAN/internal/check"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

// Write - write config to file
func Write(config models.Conf) {

	slog.Info("Writing new config to " + config.ConfPath)

	viper.SetConfigFile(config.ConfPath)
	viper.SetConfigType("yaml")

	viper.Set("HOST", config.Host)
	viper.Set("PORT", config.Port)
	viper.Set("THEME", config.Theme)
	viper.Set("COLOR", config.Color)
	viper.Set("NODEPATH", config.NodePath)
	viper.Set("LOG_LEVEL", config.LogLevel)
	viper.Set("ARP_ARGS", config.ArpArgs)
	viper.Set("ARP_STRS", config.ArpStrs)
	viper.Set("ARP_STRS_JOINED", "") // Can be set only with ENV
	viper.Set("IFACES", config.Ifaces)
	viper.Set("TIMEOUT", config.Timeout)
	viper.Set("TRIM_HIST", config.TrimHist)
	viper.Set("SHOUTRRR_URL", config.ShoutURL)

	viper.Set("USE_DB", config.UseDB)
	viper.Set("PG_CONNECT", config.PGConnect)

	viper.Set("influx_enable", config.InfluxEnable)
	viper.Set("influx_skip_tls", config.InfluxSkipTLS)
	viper.Set("influx_addr", config.InfluxAddr)
	viper.Set("influx_token", config.InfluxToken)
	viper.Set("influx_org", config.InfluxOrg)
	viper.Set("influx_bucket", config.InfluxBucket)

	viper.Set("PROMETHEUS_ENABLE", config.PrometheusEnable)

	err := viper.WriteConfig()
	check.IfError(err)
}
