package web

import (
	"net/http"

	"github.com/aceberg/WatchYourLAN/internal/conf"
	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	basePath := ""
	if conf.AppConfig.BasePath != "" {
		basePath = "/" + conf.AppConfig.BasePath
	}

	data := gin.H{
		"BasePath": basePath,
	}

	c.HTML(http.StatusOK, "index.html", data)
}
