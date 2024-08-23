package web

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/WatchYourLAN/internal/models"
)

func hostHandler(c *gin.Context) {
	var guiData models.GuiData
	guiData.Config = appConfig

	idStr := c.Param("id")
	host := getHostByID(idStr, allHosts)
	_, host.DNS = updateDNS(host)
	guiData.Host = host

	c.HTML(http.StatusOK, "header.html", guiData)
	c.HTML(http.StatusOK, "host.html", guiData)
}
