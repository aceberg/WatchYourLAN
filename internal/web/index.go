package web

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/WatchYourLAN/internal/models"
)

func indexHandler(c *gin.Context) {
	var guiData models.GuiData
	guiData.Config = appConfig
	guiData.Themes = getAllIfaces(allHosts)

	c.HTML(http.StatusOK, "index.html", guiData)
}

func historyHandler(c *gin.Context) {
	var guiData models.GuiData
	guiData.Config = appConfig

	c.HTML(http.StatusOK, "history.html", guiData)
}
