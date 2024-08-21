package web

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/WatchYourLAN/internal/arp"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

func indexHandler(c *gin.Context) {
	var guiData models.GuiData
	guiData.Config = appConfig

	allHosts = arp.Scan(appConfig.Ifaces)
	log.Println("ARPSCAN:", allHosts)

	go arp.GoScan()

	c.HTML(http.StatusOK, "header.html", guiData)
	c.HTML(http.StatusOK, "index.html", guiData)
}
