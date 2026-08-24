package web

import (
	"embed"
	"html/template"
	"log/slog"
	"net/http"

	"github.com/aceberg/WatchYourLAN/internal/api"
	"github.com/aceberg/WatchYourLAN/internal/check"
	"github.com/aceberg/WatchYourLAN/internal/conf"
	"github.com/aceberg/WatchYourLAN/internal/prometheus"
	"github.com/gin-gonic/gin"
)

// templFS - html templates
//
//go:embed templates/*
var templFS embed.FS

// pubFS - public folder
//
//go:embed public/*
var pubFS embed.FS

// Gui - start web server
func Gui() {
	const (
		colorCyan  = "\033[36m"
		colorReset = "\033[0m"
	)

	file, err := pubFS.ReadFile("public/version")
	check.IfError(err)
	conf.SetVersion(string(file)[8:])

	config := conf.GetAppConfig()
	address := config.Host + ":" + config.Port

	slog.Info(colorCyan + "\n=================================== " +
		"\n  LANventory Version: " + config.Version +
		"\n  Config dir: " + config.DirPath +
		"\n  Default DB: " + config.UseDB +
		"\n  Log level: " + config.LogLevel +
		"\n  Web GUI: http://" + address +
		"\n=================================== " + colorReset)

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())

	templ := template.Must(template.New("").ParseFS(templFS, "templates/*"))
	router.SetHTMLTemplate(templ) // templates

	router.StaticFS("/fs/", http.FS(pubFS)) // public

	router.GET("/", indexHandler)          // index.go
	router.GET("/config", indexHandler)    // index.go
	router.GET("/history", indexHandler)   // index.go
	router.GET("/activity", indexHandler)  // index.go
	router.GET("/host/*any", indexHandler) // index.go
	router.GET("/metrics", prometheus.Handler())

	api.Routes(router)

	err = router.Run(address)
	check.IfError(err)
}
