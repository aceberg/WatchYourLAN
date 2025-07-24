package api

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/WatchYourLAN/internal/check"
	"github.com/aceberg/WatchYourLAN/internal/conf"
	"github.com/aceberg/WatchYourLAN/internal/gdb"
	"github.com/aceberg/WatchYourLAN/internal/models"
	"github.com/aceberg/WatchYourLAN/internal/notify"
	"github.com/aceberg/WatchYourLAN/internal/portscan"
)

func apiAll(c *gin.Context) {

	allHosts := gdb.Select("now")

	c.IndentedJSON(http.StatusOK, allHosts)
}

func apiVersion(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, conf.AppConfig.Version)
}

func apiGetConfig(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, conf.AppConfig)
}

func apiHistory(c *gin.Context) {
	var hosts []models.Host

	mac := c.Param("mac")

	if mac != "/" {
		mac = mac[1:]
		hosts = gdb.SelectByMAC(mac)
	} else {
		hosts = gdb.Select("history")
	}

	c.IndentedJSON(http.StatusOK, hosts)
}

func apiHost(c *gin.Context) {

	idStr := c.Param("id")

	host := getHostByID(idStr) // functions.go
	_, host.DNS = check.DNS(host)

	c.IndentedJSON(http.StatusOK, host)
}

func apiHostDel(c *gin.Context) {

	idStr := c.Param("id")
	host := getHostByID(idStr) // functions.go
	gdb.Delete("now", host.ID)

	slog.Info("Deleting from DB", "host", host)

	c.IndentedJSON(http.StatusOK, "OK")
}

func apiPort(c *gin.Context) {

	addr := c.Param("addr")
	port := c.Param("port")
	state := portscan.IsOpen(addr, port)

	c.IndentedJSON(http.StatusOK, state)
}

func apiEdit(c *gin.Context) {

	idStr := c.Param("id")
	name := c.Param("name")
	toggleKnown := c.Param("known")

	host := getHostByID(idStr) // functions.go

	host.Name = name

	if toggleKnown == "/toggle" {
		host.Known = 1 - host.Known
	}

	gdb.Update("now", host)

	c.IndentedJSON(http.StatusOK, "OK")
}

func apiNotifyTest(c *gin.Context) {

	notify.Test()

	c.Status(http.StatusOK)
}

func apiStatus(c *gin.Context) {
	var status models.Stat
	var searchHosts []models.Host

	allHosts := gdb.Select("now")

	iface := c.Param("iface")
	iface = iface[1:]

	if iface != "" {
		for _, host := range allHosts {
			if iface == host.Iface {
				searchHosts = append(searchHosts, host)
			}
		}
	} else {
		searchHosts = allHosts
	}

	for _, host := range searchHosts {
		status.Total = status.Total + 1

		if host.Known > 0 {
			status.Known = status.Known + 1
		} else {
			status.Unknown = status.Unknown + 1
		}
		if host.Now > 0 {
			status.Online = status.Online + 1
		} else {
			status.Offline = status.Offline + 1
		}
	}

	c.IndentedJSON(http.StatusOK, status)
}
