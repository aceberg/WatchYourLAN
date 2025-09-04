package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", true)
}
