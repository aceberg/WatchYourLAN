package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/aceberg/WatchYourLAN/internal/models"
)

func apiAll(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, allHosts)
}

func apiHistory(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, histHosts)
}
