package api

import (
	"github.com/gin-gonic/gin"
)

// Routes - start API routes
func Routes(router *gin.Engine) {

	router.GET("/api/all", getAllHosts)                     // api.go
	router.GET("/api/config", getConfig)                    // api.go
	router.GET("/api/edit/:id/:name/*known", editHost)      // api.go
	router.GET("/api/history", getHistory)                  // api.go
	router.GET("/api/history/:mac", getHistoryByMAC)        // api.go
	router.GET("/api/history/:mac/:date", getHistoryByDate) // api.go
	router.GET("/api/host/:id", getHost)                    // api.go
	router.GET("/api/host/del/:id", delHost)                // api.go
	router.GET("/api/notify_test", notifyTest)              // api.go
	router.GET("/api/port/:addr/:port", getPortState)       // api.go
	router.GET("/api/status/*iface", getStatus)             // api.go
	router.GET("/api/version", getVersion)                  // api.go
	router.GET("/api/wol/:mac", sendWOL)                    // api.go

	router.POST("/api/config/", saveConfigHandler)                // config.go
	router.POST("/api/config_settings/", saveSettingsHandler)     // config.go
	router.POST("/api/config_influx/", saveInfluxHandler)         // config.go
	router.POST("/api/config_prometheus/", savePrometheusHandler) // config.go
}
