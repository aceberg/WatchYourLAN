package api

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/WatchYourLAN/internal/conf"
)

func saveConfigHandler(c *gin.Context) {

	conf.AppConfig.Host = c.PostForm("host")
	conf.AppConfig.Port = c.PostForm("port")
	conf.AppConfig.Theme = c.PostForm("theme")
	conf.AppConfig.Color = c.PostForm("color")
	conf.AppConfig.NodePath = c.PostForm("node")
	conf.AppConfig.ShoutURL = c.PostForm("shout")

	conf.Write(conf.AppConfig)

	slog.Info("Writing new config to " + conf.AppConfig.ConfPath)

	c.Redirect(http.StatusFound, c.Request.Referer())
}

func saveSettingsHandler(c *gin.Context) {

	conf.AppConfig.LogLevel = c.PostForm("log")
	conf.AppConfig.ArpArgs = c.PostForm("arpargs")
	conf.AppConfig.Ifaces = c.PostForm("ifaces")

	conf.AppConfig.UseDB = c.PostForm("usedb")
	conf.AppConfig.PGConnect = c.PostForm("pgconnect")

	timeout := c.PostForm("timeout")
	trimHist := c.PostForm("trim")
	conf.AppConfig.Timeout, _ = strconv.Atoi(timeout)
	conf.AppConfig.TrimHist, _ = strconv.Atoi(trimHist)

	histdb := c.PostForm("histdb")
	if histdb == "on" {
		conf.AppConfig.HistInDB = true
	} else {
		conf.AppConfig.HistInDB = false
	}

	arpStrs := c.PostFormArray("arpstrs")
	conf.AppConfig.ArpStrs = []string{}
	for _, s := range arpStrs {
		if s != "" {
			conf.AppConfig.ArpStrs = append(conf.AppConfig.ArpStrs, s)
		}
	}

	conf.Write(conf.AppConfig)

	slog.Debug("ARP_STRS", "", conf.AppConfig.ArpArgs)
	slog.Info("Writing new config to " + conf.AppConfig.ConfPath)

	// TODO:
	// updateRoutines() // routines-upd.go

	c.Redirect(http.StatusFound, c.Request.Referer())
}

func saveInfluxHandler(c *gin.Context) {

	conf.AppConfig.InfluxAddr = c.PostForm("addr")
	conf.AppConfig.InfluxToken = c.PostForm("token")
	conf.AppConfig.InfluxOrg = c.PostForm("org")
	conf.AppConfig.InfluxBucket = c.PostForm("bucket")
	enable := c.PostForm("enable")
	skip := c.PostForm("skip")

	if enable == "on" {
		conf.AppConfig.InfluxEnable = true
	} else {
		conf.AppConfig.InfluxEnable = false
	}
	if skip == "on" {
		conf.AppConfig.InfluxSkipTLS = true
	} else {
		conf.AppConfig.InfluxSkipTLS = false
	}

	conf.Write(conf.AppConfig)

	slog.Info("Writing new config to " + conf.AppConfig.ConfPath)

	c.Redirect(http.StatusFound, c.Request.Referer())
}

func savePrometheusHandler(c *gin.Context) {
	enable := c.PostForm("enable")

	conf.AppConfig.PrometheusEnable = enable == "on"

	conf.Write(conf.AppConfig)

	slog.Info("Writing new config to " + conf.AppConfig.ConfPath)

	c.Redirect(http.StatusFound, c.Request.Referer())
}
