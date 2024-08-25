package conf

import (
	"github.com/spf13/viper"

	"github.com/aceberg/WatchYourLAN/internal/check"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

// Get - get app config
func Get(path string) (config models.Conf) {

	viper.SetDefault("HOST", "0.0.0.0")
	viper.SetDefault("PORT", "8840")
	viper.SetDefault("THEME", "solar")
	viper.SetDefault("COLOR", "dark")
	viper.SetDefault("NODEPATH", "")
	viper.SetDefault("LOG_LEVEL", "info")
	viper.SetDefault("ARP_ARGS", "")
	viper.SetDefault("IFACES", "")
	viper.SetDefault("TIMEOUT", 60)
	viper.SetDefault("TRIM_HIST", 48)
	viper.SetDefault("SHOUTRRR_URL", "")
	viper.SetDefault("INFLUX_ENABLE", false)

	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	check.IfError(err)

	viper.AutomaticEnv() // Get ENVIRONMENT variables

	config.Host = viper.Get("HOST").(string)
	config.Port = viper.Get("PORT").(string)
	config.Theme = viper.Get("THEME").(string)
	config.Color = viper.Get("COLOR").(string)
	config.NodePath = viper.Get("NODEPATH").(string)
	config.LogLevel = viper.Get("LOG_LEVEL").(string)
	config.ArpArgs = viper.Get("ARP_ARGS").(string)
	config.Ifaces = viper.Get("IFACES").(string)
	config.Timeout = viper.GetInt("TIMEOUT")
	config.TrimHist = viper.GetInt("TRIM_HIST")
	config.ShoutURL = viper.Get("SHOUTRRR_URL").(string)

	config.InfluxEnable = viper.GetBool("INFLUX_ENABLE")
	config.InfluxSkipTLS = viper.GetBool("INFLUX_SKIP_TLS")
	config.InfluxAddr, _ = viper.Get("INFLUX_ADDR").(string)
	config.InfluxToken, _ = viper.Get("INFLUX_TOKEN").(string)
	config.InfluxOrg, _ = viper.Get("INFLUX_ORG").(string)
	config.InfluxBucket, _ = viper.Get("INFLUX_BUCKET").(string)

	return config
}

// Write - write config to file
func Write(config models.Conf) {

	viper.SetConfigFile(config.ConfPath)
	viper.SetConfigType("yaml")

	viper.Set("HOST", config.Host)
	viper.Set("PORT", config.Port)
	viper.Set("THEME", config.Theme)
	viper.Set("COLOR", config.Color)
	viper.Set("NODEPATH", config.NodePath)
	viper.Set("LOG_LEVEL", config.LogLevel)
	viper.Set("ARP_ARGS", config.ArpArgs)
	viper.Set("IFACES", config.Ifaces)
	viper.Set("TIMEOUT", config.Timeout)
	viper.Set("TRIM_HIST", config.TrimHist)
	viper.Set("SHOUTRRR_URL", config.ShoutURL)

	viper.Set("influx_enable", config.InfluxEnable)
	viper.Set("influx_skip_tls", config.InfluxSkipTLS)
	viper.Set("influx_addr", config.InfluxAddr)
	viper.Set("influx_token", config.InfluxToken)
	viper.Set("influx_org", config.InfluxOrg)
	viper.Set("influx_bucket", config.InfluxBucket)

	err := viper.WriteConfig()
	check.IfError(err)
}
