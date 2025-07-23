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

	slog.Info("Config dir", "path", conf.AppConfig.DirPath)

	address := conf.AppConfig.Host + ":" + conf.AppConfig.Port

	slog.Info("=================================== ")
	slog.Info("Web GUI at http://" + address)
	slog.Info("=================================== ")

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())

	templ := template.Must(template.New("").ParseFS(templFS, "templates/*"))
	router.SetHTMLTemplate(templ) // templates

	router.StaticFS("/fs/", http.FS(pubFS)) // public

	router.GET("/", indexHandler)          // index.go
	router.GET("/config", indexHandler)    // index.go
	router.GET("/history", indexHandler)   // index.go
	router.GET("/host/*any", indexHandler) // index.go
	router.GET("/metrics", prometheus.Handler(conf.AppConfig.PrometheusEnable))

	api.Routes(router)

	err := router.Run(address)
	check.IfError(err)
}
