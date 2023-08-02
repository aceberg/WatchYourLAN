package migrate

import (
	"github.com/spf13/viper"

	"github.com/aceberg/WatchYourLAN/internal/models"
)

func getOldConfig(path string) (config models.Conf) {

	viper.SetDefault("IFACE", "enp1s0")
	viper.SetDefault("DBPATH", "/data/db.sqlite")
	viper.SetDefault("GUIIP", "localhost")
	viper.SetDefault("GUIPORT", "8840")
	viper.SetDefault("TIMEOUT", "60")
	viper.SetDefault("SHOUTRRR_URL", "")
	viper.SetDefault("THEME", "solar")
	viper.SetDefault("COLOR", "light")
	viper.SetDefault("IGNOREIP", "no")
	viper.SetDefault("LOGLEVEL", "verbose")

	viper.SetConfigFile(path)
	viper.SetConfigType("env")

	viper.AutomaticEnv() // Get ENVIRONMENT variables

	err := viper.ReadInConfig()

	if err == nil {

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
	}

	return config
}
