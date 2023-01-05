package web

import (
	"net/http"
	"strconv"

	"github.com/aceberg/WatchYourLAN/internal/check"
	"github.com/aceberg/WatchYourLAN/internal/models"
	"github.com/aceberg/WatchYourLAN/internal/port"
)

func portHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig
	guiData.Icon = Icon

	ip := r.FormValue("ip")
	beginStr := r.FormValue("begin")
	endStr := r.FormValue("end")

	begin, err := strconv.Atoi(beginStr)
	check.IfError(err)

	end, err := strconv.Atoi(endStr)
	check.IfError(err)

	guiData.Themes = port.Scan(ip, "tcp", begin, end)

	execTemplate(w, "port", guiData)
}
