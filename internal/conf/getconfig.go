package conf

import (
	. "github.com/aceberg/WatchYourLAN/internal/models"
	"github.com/spf13/viper"
)

func GetConfig(path string) (config Conf) {
	viper.SetDefault("IFACE", "enp1s0")
	viper.SetDefault("DBPATH", "/data/db.sqlite")
	viper.SetDefault("GUIIP", "localhost")
	viper.SetDefault("GUIPORT", "8840")
	viper.SetDefault("TIMEOUT", "60")
	viper.SetDefault("SHOUTRRR_URL", "")
	viper.SetDefault("THEME", "solar")

	viper.SetConfigFile(path)
	viper.SetConfigType("env")
	viper.ReadInConfig()

	viper.AutomaticEnv() // Get ENVIRONMENT variables

	config.Iface = viper.Get("IFACE").(string)
	config.DbPath = viper.Get("DBPATH").(string)
	config.GuiIP = viper.Get("GUIIP").(string)
	config.GuiPort = viper.Get("GUIPORT").(string)
	config.Timeout = viper.GetInt("TIMEOUT")
	config.ShoutUrl = viper.Get("SHOUTRRR_URL").(string)
	config.Theme = viper.Get("THEME").(string)

	return config
}

func WriteConfig(path, theme string) {
	viper.SetConfigFile(path)
	viper.SetConfigType("env")
	viper.Set("THEME", theme)
	viper.WriteConfig()
}
