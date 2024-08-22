package web

import (
	// "log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/WatchYourLAN/internal/db"
)

func apiAll(c *gin.Context) {

	allHosts = db.Select(appConfig.DBPath, "now")

	c.IndentedJSON(http.StatusOK, allHosts)
}

func apiHistory(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, histHosts)
}

func apiEdit(c *gin.Context) {

	idStr := c.Param("id")
	name := c.Param("name")
	toggleKnown := c.Param("known")

	id, _ := strconv.Atoi(idStr)

	host := getHostByID(id, allHosts)
	if host.Date != "" {
		host.Name = name

		if toggleKnown == "/toggle" {
			host.Known = 1 - host.Known
		}
		// log.Println("EDIT: ", host)

		db.Update(appConfig.DBPath, "now", host)
		allHosts = db.Select(appConfig.DBPath, "now")
	}

	c.IndentedJSON(http.StatusOK, "OK")
}
