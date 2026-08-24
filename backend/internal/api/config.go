package api

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/WatchYourLAN/internal/conf"
	"github.com/aceberg/WatchYourLAN/internal/gdb"
	"github.com/aceberg/WatchYourLAN/internal/models"
	"github.com/aceberg/WatchYourLAN/internal/routines"
)

var errInvalidPositiveInt = errors.New("invalid positive integer")

var restartScanner = routines.ScanRestart

type colorRequest struct {
	Color string `json:"color"`
}

type retentionRequest struct {
	PresenceRetention     int `form:"presenceRetention" json:"presenceRetention"`
	ConnectivityRetention int `form:"connectivityRetention" json:"connectivityRetention"`
}

func saveConfigHandler(c *gin.Context) {

	_, err := conf.UpdateAppConfig(func(nextConfig *models.Conf) error {
		nextConfig.Host = c.PostForm("host")
		nextConfig.Port = c.PostForm("port")
		nextConfig.Theme = c.PostForm("theme")
		nextConfig.Color = c.PostForm("color")
		nextConfig.NodePath = c.PostForm("node")
		nextConfig.ShoutURL = applySecretUpdate(nextConfig.ShoutURL, c.PostForm("shout"), c.PostForm("clear_shout"))
		return nil
	})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to write config"})
		return
	}

	c.Redirect(http.StatusFound, c.Request.Referer())
}

func saveColorHandler(c *gin.Context) {
	color := c.PostForm("color")

	if color == "" {
		var req colorRequest
		if err := c.ShouldBindJSON(&req); err == nil {
			color = req.Color
		}
	}

	if color != "dark" && color != "light" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid color"})
		return
	}

	nextConfig, err := conf.UpdateAppConfig(func(nextConfig *models.Conf) error {
		nextConfig.Color = color
		return nil
	})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to write config"})
		return
	}

	c.IndentedJSON(http.StatusOK, toPublicConfig(nextConfig))
}

func saveRetentionHandler(c *gin.Context) {
	var req retentionRequest
	if err := c.ShouldBind(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid retention"})
		return
	}

	if req.PresenceRetention < 1 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid presenceRetention"})
		return
	}
	if req.ConnectivityRetention < 1 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid connectivityRetention"})
		return
	}

	nextConfig, err := conf.UpdateAppConfig(func(nextConfig *models.Conf) error {
		nextConfig.TrimHist = req.PresenceRetention
		nextConfig.ConnectivityRetention = req.ConnectivityRetention
		return nil
	})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to write config"})
		return
	}

	c.IndentedJSON(http.StatusOK, toPublicConfig(nextConfig))
}

func saveSettingsHandler(c *gin.Context) {

	currentConfig := conf.GetAppConfig()
	nextConfig := currentConfig
	nextConfig.LogLevel = c.PostForm("log")
	nextConfig.ArpArgs = c.PostForm("arpargs")
	nextConfig.Ifaces = c.PostForm("ifaces")

	useDB := c.PostForm("usedb")
	if useDB != "sqlite" && useDB != "postgres" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid usedb"})
		return
	}
	nextConfig.UseDB = useDB
	nextConfig.PGConnect = applySecretUpdate(nextConfig.PGConnect, c.PostForm("pgconnect"), c.PostForm("clear_pgconnect"))

	if !isValidLogLevel(nextConfig.LogLevel) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid log"})
		return
	}

	timeout, err := parsePositiveInt(c.PostForm("timeout"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid timeout"})
		return
	}

	trimHist := nextConfig.TrimHist
	if rawTrimHist := c.PostForm("trim"); rawTrimHist != "" {
		trimHist, err = parsePositiveInt(rawTrimHist)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid trim"})
			return
		}
	}

	connectivityRetention := nextConfig.ConnectivityRetention
	if connectivityRetention < 1 {
		connectivityRetention = trimHist
	}
	if rawConnectivityRetention := c.PostForm("connectivity_retention"); rawConnectivityRetention != "" {
		connectivityRetention, err = parsePositiveInt(rawConnectivityRetention)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid connectivity_retention"})
			return
		}
	}

	nextConfig.Timeout = timeout
	nextConfig.TrimHist = trimHist
	nextConfig.ConnectivityRetention = connectivityRetention

	arpStrs := c.PostFormArray("arpstrs")
	nextConfig.ArpStrs = []string{}
	for _, s := range arpStrs {
		if s != "" {
			nextConfig.ArpStrs = append(nextConfig.ArpStrs, s)
		}
	}

	dbChanged := nextConfig.UseDB != currentConfig.UseDB || nextConfig.PGConnect != currentConfig.PGConnect
	if dbChanged {
		if err := gdb.Reconnect(nextConfig); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to reconnect database"})
			return
		}
	}

	_, err = conf.UpdateAppConfig(func(config *models.Conf) error {
		config.LogLevel = nextConfig.LogLevel
		config.ArpArgs = nextConfig.ArpArgs
		config.Ifaces = nextConfig.Ifaces
		config.UseDB = nextConfig.UseDB
		config.PGConnect = nextConfig.PGConnect
		config.Timeout = nextConfig.Timeout
		config.TrimHist = nextConfig.TrimHist
		config.ConnectivityRetention = nextConfig.ConnectivityRetention
		config.ArpStrs = append([]string(nil), nextConfig.ArpStrs...)
		return nil
	})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to write config"})
		return
	}

	restartScanner()

	c.Redirect(http.StatusFound, c.Request.Referer())
}

func parsePositiveInt(value string) (int, error) {
	parsed, err := strconv.Atoi(value)
	if err != nil || parsed < 1 {
		return 0, errInvalidPositiveInt
	}

	return parsed, nil
}

func isValidLogLevel(value string) bool {
	switch value {
	case "debug", "info", "warn", "error":
		return true
	default:
		return false
	}
}

func saveInfluxHandler(c *gin.Context) {

	enable := c.PostForm("enable")
	skip := c.PostForm("skip")
	_, err := conf.UpdateAppConfig(func(nextConfig *models.Conf) error {
		nextConfig.InfluxAddr = c.PostForm("addr")
		nextConfig.InfluxToken = applySecretUpdate(nextConfig.InfluxToken, c.PostForm("token"), c.PostForm("clear_influx_token"))
		nextConfig.InfluxOrg = c.PostForm("org")
		nextConfig.InfluxBucket = c.PostForm("bucket")
		nextConfig.InfluxEnable = enable == "on"
		nextConfig.InfluxSkipTLS = skip == "on"
		return nil
	})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to write config"})
		return
	}

	c.Redirect(http.StatusFound, c.Request.Referer())
}

func savePrometheusHandler(c *gin.Context) {
	enable := c.PostForm("enable")

	_, err := conf.UpdateAppConfig(func(nextConfig *models.Conf) error {
		nextConfig.PrometheusEnable = enable == "on"
		return nil
	})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to write config"})
		return
	}

	c.Redirect(http.StatusFound, c.Request.Referer())
}

func applySecretUpdate(currentValue, submittedValue, clearValue string) string {
	if isTruthyFormValue(clearValue) {
		return ""
	}
	if submittedValue != "" {
		return submittedValue
	}

	return currentValue
}

func isTruthyFormValue(value string) bool {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "1", "on", "true", "yes":
		return true
	default:
		return false
	}
}
