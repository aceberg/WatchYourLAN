package web

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/WatchYourLAN/internal/db"
	"github.com/aceberg/WatchYourLAN/internal/models"
	"github.com/aceberg/WatchYourLAN/internal/portscan"
)

func apiAll(c *gin.Context) {

	allHosts = db.Select("now")

	c.IndentedJSON(http.StatusOK, allHosts)
}

func apiHistory(c *gin.Context) {
	var hosts []models.Host

	if appConfig.HistInDB {
		histHosts = db.Select("history")
	}

	mac := c.Param("mac")

	if mac != "/" {
		mac = mac[1:]
		hosts = getHostsByMAC(mac, histHosts)
	} else {
		hosts = histHosts
	}

	c.IndentedJSON(http.StatusOK, hosts)
}

func apiHost(c *gin.Context) {

	idStr := c.Param("id")

	host := getHostByID(idStr, allHosts) // functions.go

	c.IndentedJSON(http.StatusOK, host)
}

func apiHostDel(c *gin.Context) {

	idStr := c.Param("id")
	host := getHostByID(idStr, allHosts) // functions.go
	db.Delete("now", host.ID)
	allHosts = db.Select("now")

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

	host := getHostByID(idStr, allHosts) // functions.go
	if host.Date != "" {
		host.Name = name

		if toggleKnown == "/toggle" {
			host.Known = 1 - host.Known
		}
		// log.Println("EDIT: ", host)

		db.Update("now", host)
		allHosts = db.Select("now")
	}

	c.IndentedJSON(http.StatusOK, "OK")
}
