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

	viper.Set("HOST", stripQuotes(config.Host))
	viper.Set("PORT", stripQuotes(config.Port))
	viper.Set("THEME", stripQuotes(config.Theme))
	viper.Set("COLOR", stripQuotes(config.Color))
	viper.Set("NODEPATH", stripQuotes(config.NodePath))
	viper.Set("LOG_LEVEL", stripQuotes(config.LogLevel))
	viper.Set("ARP_ARGS", stripQuotes(config.ArpArgs))
	viper.Set("ARP_STRS", config.ArpStrs)
	viper.Set("ARP_STRS_JOINED", "") // Can be set only with ENV
	viper.Set("IFACES", stripQuotes(config.Ifaces))
	viper.Set("TIMEOUT", config.Timeout)
	viper.Set("TRIM_HIST", config.TrimHist)
	viper.Set("SHOUTRRR_URL", stripQuotes(config.ShoutURL))

	viper.Set("USE_DB", stripQuotes(config.UseDB))
	viper.Set("PG_CONNECT", stripQuotes(config.PGConnect))

	viper.Set("influx_enable", config.InfluxEnable)
	viper.Set("influx_skip_tls", config.InfluxSkipTLS)
	viper.Set("influx_addr", stripQuotes(config.InfluxAddr))
	viper.Set("influx_token", stripQuotes(config.InfluxToken))
	viper.Set("influx_org", stripQuotes(config.InfluxOrg))
	viper.Set("influx_bucket", stripQuotes(config.InfluxBucket))

	viper.Set("PROMETHEUS_ENABLE", config.PrometheusEnable)

	err := viper.WriteConfig()
	check.IfError(err)
}
