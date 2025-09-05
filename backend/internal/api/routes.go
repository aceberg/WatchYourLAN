package api

import (
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Routes - start API routes
func Routes(router *gin.Engine) {

	r0 := router.Group("/api")
	{
		r0.GET("/all", getAllHosts)                // api-hosts.go
		r0.GET("/edit/:id/:name/*known", editHost) // api-hosts.go
		r0.GET("/host/:id", getHost)               // api-hosts.go
		r0.GET("/host/del/:id", delHost)           // api-hosts.go

		r0.GET("/config", getConfig)        // api-system.go
		r0.GET("/notify_test", notifyTest)  // api-system.go
		r0.GET("/status/*iface", getStatus) // api-system.go
		r0.GET("/version", getVersion)      // api-system.go

		r0.GET("/history", getHistory)                  // api-history.go
		r0.GET("/history/:mac", getHistoryByMAC)        // api-history.go
		r0.GET("/history/:mac/:date", getHistoryByDate) // api-history.go

		r0.GET("/port/:addr/:port", getPortState) // api-network.go
		r0.GET("/wol/:mac", sendWOL)              // api-network.go

		r0.POST("/config/", saveConfigHandler)                // config.go
		r0.POST("/config_settings/", saveSettingsHandler)     // config.go
		r0.POST("/config_influx/", saveInfluxHandler)         // config.go
		r0.POST("/config_prometheus/", savePrometheusHandler) // config.go
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
