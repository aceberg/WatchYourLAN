package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/WatchYourLAN/internal/conf"
	"github.com/aceberg/WatchYourLAN/internal/gdb"
	"github.com/aceberg/WatchYourLAN/internal/routines"
)

func saveConfigHandler(c *gin.Context) {

	conf.AppConfig.Host = c.PostForm("host")
	conf.AppConfig.Port = c.PostForm("port")
	conf.AppConfig.Theme = c.PostForm("theme")
	conf.AppConfig.Color = c.PostForm("color")
	conf.AppConfig.NodePath = c.PostForm("node")
	conf.AppConfig.ShoutURL = c.PostForm("shout")

	conf.Write(conf.AppConfig)

	c.Redirect(http.StatusFound, c.Request.Referer())
}

func saveSettingsHandler(c *gin.Context) {

	conf.AppConfig.LogLevel = c.PostForm("log")
	conf.AppConfig.ArpArgs = c.PostForm("arpargs")
	conf.AppConfig.Ifaces = c.PostForm("ifaces")

	useDB := c.PostForm("usedb")
	pgConnect := c.PostForm("pgconnect")

	if useDB != conf.AppConfig.UseDB || pgConnect != conf.AppConfig.PGConnect {
		conf.AppConfig.UseDB = c.PostForm("usedb")
		conf.AppConfig.PGConnect = c.PostForm("pgconnect")
		gdb.Connect()
	}

	timeout := c.PostForm("timeout")
	trimHist := c.PostForm("trim")
	conf.AppConfig.Timeout, _ = strconv.Atoi(timeout)
	conf.AppConfig.TrimHist, _ = strconv.Atoi(trimHist)

	arpStrs := c.PostFormArray("arpstrs")
	conf.AppConfig.ArpStrs = []string{}
	for _, s := range arpStrs {
		if s != "" {
			conf.AppConfig.ArpStrs = append(conf.AppConfig.ArpStrs, s)
		}
	}

	conf.Write(conf.AppConfig)

	routines.ScanRestart()

	c.Redirect(http.StatusFound, c.Request.Referer())
}

func saveInfluxHandler(c *gin.Context) {

	conf.AppConfig.InfluxAddr = c.PostForm("addr")
	conf.AppConfig.InfluxToken = c.PostForm("token")
	conf.AppConfig.InfluxOrg = c.PostForm("org")
	conf.AppConfig.InfluxBucket = c.PostForm("bucket")

	enable := c.PostForm("enable")
	skip := c.PostForm("skip")
	conf.AppConfig.InfluxEnable = enable == "on"
	conf.AppConfig.InfluxSkipTLS = skip == "on"

	conf.Write(conf.AppConfig)

	c.Redirect(http.StatusFound, c.Request.Referer())
}

func savePrometheusHandler(c *gin.Context) {
	enable := c.PostForm("enable")

	conf.AppConfig.PrometheusEnable = enable == "on"

	conf.Write(conf.AppConfig)

	c.Redirect(http.StatusFound, c.Request.Referer())
}
