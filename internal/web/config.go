package web

import (
	"log/slog"
	"net/http"
	// "strconv"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/WatchYourLAN/internal/check"
	// "github.com/aceberg/WatchYourLAN/internal/conf"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

func configHandler(c *gin.Context) {
	var guiData models.GuiData

	guiData.Config = appConfig

	guiData.Themes = []string{"cerulean", "cosmo", "cyborg", "darkly", "emerald", "flatly", "grass", "grayscale", "journal", "litera", "lumen", "lux", "materia", "minty", "morph", "ocean", "pulse", "quartz", "sand", "sandstone", "simplex", "sketchy", "slate", "solar", "spacelab", "superhero", "united", "vapor", "wood", "yeti", "zephyr"}

	file, err := pubFS.ReadFile("public/version")
	check.IfError(err)
	version := string(file)
	guiData.Version = version[8:]

	c.HTML(http.StatusOK, "header.html", guiData)
	c.HTML(http.StatusOK, "config.html", guiData)
}

func saveConfigHandler(c *gin.Context) {

	// appConfig.Host = c.PostForm("host")
	// appConfig.Port = c.PostForm("port")
	// appConfig.Theme = c.PostForm("theme")
	// appConfig.Color = c.PostForm("color")
	// tStr := c.PostForm("timeout")
	// hStr := c.PostForm("trim")

	// if tStr != "" {
	// 	appConfig.Timeout, _ = strconv.Atoi(tStr)
	// }
	// if hStr != "" {
	// 	appConfig.HistTrim, _ = strconv.Atoi(hStr)
	// }

	// conf.Write(appConfig)

	slog.Info("writing new config to " + appConfig.ConfPath)

	c.Redirect(http.StatusFound, "/config")
}
