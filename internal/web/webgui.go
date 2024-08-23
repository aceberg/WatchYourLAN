package web

import (
	"html/template"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/WatchYourLAN/internal/check"
	"github.com/aceberg/WatchYourLAN/internal/conf"
)

// Gui - start web server
func Gui(dirPath, nodePath string) {

	confPath := dirPath + "/config.yaml"
	check.Path(confPath)

	appConfig = conf.Get(confPath)

	appConfig.DirPath = dirPath
	appConfig.ConfPath = confPath
	appConfig.DBPath = dirPath + "/scan.db"
	appConfig.NodePath = nodePath

	updateScan() // scan.go

	slog.Info("config", "path", appConfig.DirPath)

	address := appConfig.Host + ":" + appConfig.Port

	slog.Info("=================================== ")
	slog.Info("Web GUI at http://" + address)
	slog.Info("=================================== ")

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	templ := template.Must(template.New("").ParseFS(templFS, "templates/*"))
	router.SetHTMLTemplate(templ) // templates

	router.StaticFS("/fs/", http.FS(pubFS)) // public

	router.GET("/api/all", apiAll)                    // api.go
	router.GET("/api/edit/:id/:name/*known", apiEdit) // api.go
	router.GET("/api/history", apiHistory)            // api.go
	router.GET("/api/host", apiHost)                  // api.go
	router.GET("/api/port/:addr/:port", apiPort)      // api.go

	router.GET("/", indexHandler)           // index.go
	router.GET("/history/", historyHandler) // index.go
	router.GET("/host/:id", hostHandler)    // host.go
	router.GET("/config/", configHandler)   // config.go

	router.POST("/config/", saveConfigHandler) // config.go

	err := router.Run(address)
	check.IfError(err)
}
