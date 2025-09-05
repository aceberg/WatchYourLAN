package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/WatchYourLAN/internal/gdb"
)

// getHistory godoc
// @Summary      Get full history
// @Description  Retrieve the complete history of all hosts. Not recommended, the output can be a lot
// @Tags         history
// @Produce      json
// @Success      200  {array}   models.Host
// @Router       /history [get]
func getHistory(c *gin.Context) {
	hosts, _ := gdb.Select("history")
	c.IndentedJSON(http.StatusOK, hosts)
}

// getHistoryByMAC godoc
// @Summary      Get history by MAC
// @Description  Retrieve the latest history entries for a specific host by MAC address
// @Tags         history
// @Produce      json
// @Param        mac   path      string  true "MAC address of the host"
// @Param        num   query     int     true "Number of history entries to return"
// @Success      200   {array}   models.Host
// @Router       /history/{mac} [get]
func getHistoryByMAC(c *gin.Context) {
	mac := c.Param("mac")
	numStr := c.Query("num")
	num, _ := strconv.Atoi(numStr)
	hosts := gdb.SelectLatest(mac, num)
	c.IndentedJSON(http.StatusOK, hosts)
}

// getHistoryByDate godoc
// @Summary      Get history by date
// @Description  Retrieve history for a specific host on a given date
// @Description  The date format is flexible and can be:
// @Description  - Year only: `2025`
// @Description  - Year + month: `2025-09`
// @Description  - Full date: `2025-09-06`
// @Description  - Full timestamp: `2025-09-06 00:58:26`
// @Tags         history
// @Produce      json
// @Param        mac   path      string  true  "MAC address of the host"
// @Param        date  path      string  true  "Date filter (supports YYYY, YYYY-MM, YYYY-MM-DD, YYYY-MM-DD HH:mm:ss)"
// @Success      200   {array}   models.Host
// @Router       /history/{mac}/{date} [get]
func getHistoryByDate(c *gin.Context) {
	mac := c.Param("mac")
	date := c.Param("date")
	hosts := gdb.SelectByDate(mac, date)
	c.IndentedJSON(http.StatusOK, hosts)
}
