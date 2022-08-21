package main

import (
    // "fmt"
	"github.com/spf13/viper"
)

func get_config(path string) (config Conf) {
	viper.SetDefault("IFACE", "eth0")
	viper.SetDefault("DBPATH", "/data/db.sqlite")
	viper.SetDefault("GUIIP", "localhost")
	viper.SetDefault("GUIPORT", "8840")
	viper.SetDefault("TIMEOUT", "60")

    viper.SetConfigFile(path)
	viper.SetConfigType("env")
    viper.ReadInConfig()

	viper.AutomaticEnv() // Get ENVIRONMENT variables

	config.Iface = viper.Get("IFACE").(string)
	config.DbPath = viper.Get("DBPATH").(string)
	config.GuiIP = viper.Get("GUIIP").(string)
	config.GuiPort = viper.Get("GUIPORT").(string)
	config.Timeout = viper.GetInt("TIMEOUT")

    // fmt.Println(viper.Get("DBPATH"))

	return config
}