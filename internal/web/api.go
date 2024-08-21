package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func apiAll(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, allHosts)
}