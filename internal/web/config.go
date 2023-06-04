package web

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/tidwall/gjson"

	"github.com/aceberg/WatchYourLAN/internal/check"
	"github.com/aceberg/WatchYourLAN/internal/conf"
	"github.com/aceberg/WatchYourLAN/internal/db"
	"github.com/aceberg/WatchYourLAN/internal/models"
	"github.com/aceberg/WatchYourLAN/internal/notify"
	"github.com/aceberg/WatchYourLAN/internal/scan"
)

func configHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig

	file, err := TemplHTML.ReadFile(TemplPath + "version")
	check.IfError(err)

	version := string(file)
	guiData.Version = version[8:]

	url := "https://api.github.com/repos/aceberg/WatchYourLAN/releases/latest"
	req, err := http.Get(url)
	check.IfError(err)

	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	check.IfError(err)

	bodyStr := string(body)
	release := gjson.Get(bodyStr, "tag_name")
	guiData.Release = release.String()

	guiData.Themes = []string{"cerulean", "cosmo", "cyborg", "darkly", "flatly", "journal", "litera", "lumen", "lux", "materia", "minty", "morph", "pulse", "quartz", "sandstone", "simplex", "sketchy", "slate", "solar", "spacelab", "superhero", "united", "vapor", "yeti", "zephyr"}

	execTemplate(w, "config", guiData)
}

func saveConfigHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	AppConfig.Iface = r.FormValue("iface")
	AppConfig.DbPath = r.FormValue("dbpath")
	AppConfig.GuiIP = r.FormValue("host")
	AppConfig.GuiPort = r.FormValue("port")
	AppConfig.ShoutURL = r.FormValue("shout")
	AppConfig.Theme = r.FormValue("theme")
	AppConfig.Color = r.FormValue("color")
	AppConfig.IgnoreIP = r.FormValue("ignoreip")
	AppConfig.LogLevel = r.FormValue("loglevel")

	timeout := r.FormValue("timeout")
	AppConfig.Timeout, err = strconv.Atoi(timeout)
	check.IfError(err)

	close(QuitScan)

	conf.Write(ConfigPath, AppConfig)

	log.Println("INFO: writing new config")

	QuitScan = make(chan bool)
	go scan.Start(AppConfig, QuitScan)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}

func clearHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("INFO: delting all hosts from DB")

	db.Clear(AppConfig.DbPath)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}

func testNotifyHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("INFO: Test notification send")

	msg := "Test notification"
	notify.Shoutrrr(msg, AppConfig.ShoutURL)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}
