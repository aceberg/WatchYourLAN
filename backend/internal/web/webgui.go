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
	conf.AppConfig.Version = string(file)[8:]

	address := conf.AppConfig.Host + ":" + conf.AppConfig.Port

	slog.Info(colorCyan + "\n=================================== " +
		"\n  WatchYourLAN Version: " + conf.AppConfig.Version +
		"\n  Config dir: " + conf.AppConfig.DirPath +
		"\n  Default DB: " + conf.AppConfig.UseDB +
		"\n  Log level: " + conf.AppConfig.LogLevel +
		"\n  Web GUI: http://" + address + "/" + conf.AppConfig.BasePath +
		"\n=================================== " + colorReset)

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())

	templ := template.Must(template.New("").ParseFS(templFS, "templates/*"))
	router.SetHTMLTemplate(templ) // templates

	var routerGroup gin.IRouter = router

	if conf.AppConfig.BasePath != "" {
		routerGroup = router.Group(conf.AppConfig.BasePath)
	}

    routerGroup.StaticFS("/fs/", http.FS(pubFS))

	routerGroup.GET("/", indexHandler)
	routerGroup.GET("/config", indexHandler)
	routerGroup.GET("/history", indexHandler)
	routerGroup.GET("/host/*any", indexHandler)
	routerGroup.GET("/metrics", prometheus.Handler())

	api.Routes(routerGroup)

	err = router.Run(address)
	check.IfError(err)
}
