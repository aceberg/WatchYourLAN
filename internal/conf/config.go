package conf

import (
	"github.com/spf13/viper"

	"github.com/aceberg/WatchYourLAN/internal/auth"
	"github.com/aceberg/WatchYourLAN/internal/check"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

// Get - get app config
func Get(path string) (config models.Conf, authConf auth.Conf) {

	viper.SetDefault("IFACE", "enp1s0")
	viper.SetDefault("DBPATH", "/data/db.sqlite")
	viper.SetDefault("GUIIP", "0.0.0.0")
	viper.SetDefault("GUIPORT", "8840")
	viper.SetDefault("TIMEOUT", "60")
	viper.SetDefault("SHOUTRRR_URL", "")
	viper.SetDefault("THEME", "solar")
	viper.SetDefault("COLOR", "light")
	viper.SetDefault("IGNOREIP", "no")
	viper.SetDefault("LOGLEVEL", "verbose")
	viper.SetDefault("AUTH_USER", "")
	viper.SetDefault("AUTH_PASSWORD", "")
	viper.SetDefault("AUTH_EXPIRE", "7d")
	viper.SetDefault("HISTORY_DAYS", "30")

	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	check.IfError(err)

	viper.AutomaticEnv() // Get ENVIRONMENT variables

	config.Iface = viper.Get("IFACE").(string)
	config.DbPath = viper.Get("DBPATH").(string)
	config.GuiIP = viper.Get("GUIIP").(string)
	config.GuiPort = viper.Get("GUIPORT").(string)
	config.Timeout = viper.GetInt("TIMEOUT")
	config.ShoutURL = viper.Get("SHOUTRRR_URL").(string)
	config.Theme = viper.Get("THEME").(string)
	config.Color = viper.Get("COLOR").(string)
	config.IgnoreIP = viper.Get("IGNOREIP").(string)
	config.LogLevel = viper.Get("LOGLEVEL").(string)
	config.HistDays = viper.Get("HISTORY_DAYS").(string)
	authConf.Auth = viper.GetBool("AUTH")
	authConf.User, _ = viper.Get("AUTH_USER").(string)
	authConf.Password, _ = viper.Get("AUTH_PASSWORD").(string)
	authConf.ExpStr, _ = viper.Get("AUTH_EXPIRE").(string)

	authConf.Expire = auth.ToTime(authConf.ExpStr)
	config.Auth = authConf.Auth

	return config, authConf
}

// Write - write config to file
func Write(path string, config models.Conf, authConf auth.Conf) {

	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")

	viper.Set("IFACE", config.Iface)
	viper.Set("DBPATH", config.DbPath)
	viper.Set("GUIIP", config.GuiIP)
	viper.Set("GUIPORT", config.GuiPort)
	viper.Set("TIMEOUT", config.Timeout)
	viper.Set("SHOUTRRR_URL", config.ShoutURL)
	viper.Set("THEME", config.Theme)
	viper.Set("COLOR", config.Color)
	viper.Set("IGNOREIP", config.IgnoreIP)
	viper.Set("LOGLEVEL", config.LogLevel)
	viper.Set("HISTORY_DAYS", config.HistDays)

	viper.Set("auth", authConf.Auth)
	viper.Set("auth_user", authConf.User)
	viper.Set("auth_password", authConf.Password)
	viper.Set("auth_expire", authConf.ExpStr)

	err := viper.WriteConfig()
	check.IfError(err)
}
