package web

import (
	"html/template"
	"log/slog"
	"net/http"

	"github.com/aceberg/WatchYourLAN/internal/check"
	"github.com/aceberg/WatchYourLAN/internal/conf"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Gui - start web server
func Gui(dirPath, nodePath string) {

	confPath := dirPath + "/config_v2.yaml"
	check.Path(confPath)

	appConfig = conf.Get(confPath)

	appConfig.DirPath = dirPath
	appConfig.ConfPath = confPath
	appConfig.DBPath = dirPath + "/scan.db"
	appConfig.NodePath = nodePath

	quitScan = make(chan bool)
	updateRoutines()        // routines-upd.go
	go trimHistoryRoutine() // trim-history.go

	slog.Info("Config dir", "path", appConfig.DirPath)

	address := appConfig.Host + ":" + appConfig.Port

	slog.Info("=================================== ")
	slog.Info("Web GUI at http://" + address)
	slog.Info("=================================== ")

	gin.SetMode(gin.ReleaseMode)
	// router := gin.Default()
	router := gin.New()
	router.Use(gin.Recovery())

	templ := template.Must(template.New("").ParseFS(templFS, "templates/*"))
	router.SetHTMLTemplate(templ) // templates

	router.StaticFS("/fs/", http.FS(pubFS)) // public

	router.GET("/api/all", apiAll)                    // api.go
	router.GET("/api/edit/:id/:name/*known", apiEdit) // api.go
	router.GET("/api/history/*mac", apiHistory)       // api.go
	router.GET("/api/host/:id", apiHost)              // api.go
	router.GET("/api/host/del/:id", apiHostDel)       // api.go
	router.GET("/api/notify_test", apiNotifyTest)     // api.go
	router.GET("/api/port/:addr/:port", apiPort)      // api.go
	router.GET("/api/status/*iface", apiStatus)       // api.go

	router.GET("/", indexHandler)           // index.go
	router.GET("/history/", historyHandler) // index.go
	router.GET("/host/:id", hostHandler)    // host.go
	router.GET("/config/", configHandler)   // config.go

	router.POST("/config/", saveConfigHandler)            // config.go
	router.POST("/config_settings/", saveSettingsHandler) // config.go
	router.POST("/config_influx/", saveInfluxHandler)     // config.go

	if appConfig.PrometheusEnable {
		router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	}

	err := router.Run(address)
	check.IfError(err)
}
