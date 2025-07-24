package api

import (
	"github.com/gin-gonic/gin"
)

// Routes - start API routes
func Routes(router *gin.Engine) {

	router.GET("/api/all", apiAll)                    // api.go
	router.GET("/api/config", apiGetConfig)           // api.go
	router.GET("/api/edit/:id/:name/*known", apiEdit) // api.go
	router.GET("/api/history/*mac", apiHistory)       // api.go
	router.GET("/api/host/:id", apiHost)              // api.go
	router.GET("/api/host/del/:id", apiHostDel)       // api.go
	router.GET("/api/notify_test", apiNotifyTest)     // api.go
	router.GET("/api/port/:addr/:port", apiPort)      // api.go
	router.GET("/api/status/*iface", apiStatus)       // api.go
	router.GET("/api/version", apiVersion)            // api.go

	router.POST("/api/config/", saveConfigHandler)                // config.go
	router.POST("/api/config_settings/", saveSettingsHandler)     // config.go
	router.POST("/api/config_influx/", saveInfluxHandler)         // config.go
	router.POST("/api/config_prometheus/", savePrometheusHandler) // config.go
}
