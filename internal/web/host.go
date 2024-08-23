package web

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/WatchYourLAN/internal/models"
)

func hostHandler(c *gin.Context) {
	var guiData models.GuiData
	guiData.Config = appConfig

	idStr := c.Param("id")
	log.Println("ID =", idStr)

	guiData.Version = idStr

	c.HTML(http.StatusOK, "header.html", guiData)
	c.HTML(http.StatusOK, "host.html", guiData)
}
