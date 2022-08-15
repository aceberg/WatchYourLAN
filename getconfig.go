package main

import (
    // "fmt"
	"github.com/spf13/viper"
)

func get_config(path string) (config Conf) {
    viper.SetConfigFile(path)
	viper.SetConfigType("env")
    viper.ReadInConfig()

	viper.AutomaticEnv() // Get ENVIRONMENT variables

	config.Iface = viper.Get("IFACE").(string)
	config.DbPath = viper.Get("DBPATH").(string)
	config.GuiIP = viper.Get("GUIIP").(string)
	config.GuiPort = viper.Get("GUIPORT").(string)

    // fmt.Println(viper.Get("GUIPORT"))

	return config
}