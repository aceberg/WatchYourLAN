package web

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/WatchYourLAN/internal/conf"
)

func saveConfigHandler(c *gin.Context) {

	appConfig.Host = c.PostForm("host")
	appConfig.Port = c.PostForm("port")
	appConfig.Theme = c.PostForm("theme")
	appConfig.Color = c.PostForm("color")
	appConfig.NodePath = c.PostForm("node")
	appConfig.ShoutURL = c.PostForm("shout")

	conf.Write(appConfig)

	slog.Info("Writing new config to " + appConfig.ConfPath)

	c.Redirect(http.StatusFound, c.Request.Referer())
}

func saveSettingsHandler(c *gin.Context) {

	appConfig.LogLevel = c.PostForm("log")
	appConfig.ArpArgs = c.PostForm("arpargs")
	appConfig.Ifaces = c.PostForm("ifaces")

	appConfig.UseDB = c.PostForm("usedb")
	appConfig.PGConnect = c.PostForm("pgconnect")

	timeout := c.PostForm("timeout")
	trimHist := c.PostForm("trim")
	appConfig.Timeout, _ = strconv.Atoi(timeout)
	appConfig.TrimHist, _ = strconv.Atoi(trimHist)

	histdb := c.PostForm("histdb")
	if histdb == "on" {
		appConfig.HistInDB = true
	} else {
		appConfig.HistInDB = false
	}

	arpStrs := c.PostFormArray("arpstrs")
	appConfig.ArpStrs = []string{}
	for _, s := range arpStrs {
		if s != "" {
			appConfig.ArpStrs = append(appConfig.ArpStrs, s)
		}
	}

	conf.Write(appConfig)

	slog.Debug("ARP_STRS", "", appConfig.ArpArgs)
	slog.Info("Writing new config to " + appConfig.ConfPath)

	updateRoutines() // routines-upd.go

	c.Redirect(http.StatusFound, c.Request.Referer())
}

func saveInfluxHandler(c *gin.Context) {

	appConfig.InfluxAddr = c.PostForm("addr")
	appConfig.InfluxToken = c.PostForm("token")
	appConfig.InfluxOrg = c.PostForm("org")
	appConfig.InfluxBucket = c.PostForm("bucket")
	enable := c.PostForm("enable")
	skip := c.PostForm("skip")

	if enable == "on" {
		appConfig.InfluxEnable = true
	} else {
		appConfig.InfluxEnable = false
	}
	if skip == "on" {
		appConfig.InfluxSkipTLS = true
	} else {
		appConfig.InfluxSkipTLS = false
	}

	conf.Write(appConfig)

	slog.Info("Writing new config to " + appConfig.ConfPath)

	c.Redirect(http.StatusFound, c.Request.Referer())
}

func savePrometheusHandler(c *gin.Context) {
	enable := c.PostForm("enable")

	appConfig.PrometheusEnable = enable == "on"

	conf.Write(appConfig)

	slog.Info("Writing new config to " + appConfig.ConfPath)

	c.Redirect(http.StatusFound, c.Request.Referer())
}
