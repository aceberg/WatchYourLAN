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
	viper.SetDefault("SCANER", "arpscan")
	viper.SetDefault("ARPARGS", "")
	viper.SetDefault("IFACES", "")
	viper.SetDefault("TIMEOUT", 60)
	viper.SetDefault("IGNOREIP", "no")

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
	config.Scaner = viper.Get("SCANER").(string)
	config.ArpArgs = viper.Get("ARPARGS").(string)
	config.Ifaces = viper.Get("IFACES").(string)
	config.Timeout = viper.GetInt("TIMEOUT")
	config.IgnoreIP = viper.Get("IGNOREIP").(string)

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
	viper.Set("SCANER", config.Scaner)
	viper.Set("ARPARGS", config.ArpArgs)
	viper.Set("IFACES", config.Ifaces)
	viper.Set("TIMEOUT", config.Timeout)
	viper.Set("IGNOREIP", config.IgnoreIP)

	err := viper.WriteConfig()
	check.IfError(err)
}
