package api

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/aceberg/WatchYourLAN/internal/check"
	"github.com/aceberg/WatchYourLAN/internal/gdb"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

// getAllHosts godoc
// @Summary      Get all hosts
// @Description  Retrieve all hosts from the database
// @Tags         hosts
// @Produce      json
// @Success      200  {array}   models.Host
// @Router       /all [get]
func getAllHosts(c *gin.Context) {
	allHosts, _ := gdb.Select("now")
	c.IndentedJSON(http.StatusOK, allHosts)
}

// getHost godoc
// @Summary      Get host by ID
// @Description  Retrieve detailed information about a host by its unique ID
// @Tags         hosts
// @Produce      json
// @Param        id   path      string  true  "Host ID"
// @Success      200  {object}  models.Host
// @Router       /host/{id} [get]
func getHost(c *gin.Context) {
	idStr := c.Param("id")
	host, err := getHostByID(idStr) // functions.go
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, host.DNS = check.DNS(host)
	c.IndentedJSON(http.StatusOK, host)
}

// setHostDeviceType godoc
// @Summary      Set host device type
// @Description  Update only a host's manually assigned device type
// @Tags         hosts
// @Accept       json
// @Produce      json
// @Param        id    path      string  true  "Host ID"
// @Param        body  body      object  true  "Device type payload"
// @Success      200   {object}  models.Host
// @Router       /host/{id}/type [patch]
func setHostDeviceType(c *gin.Context) {
	idStr := c.Param("id")
	host, err := getHostByID(idStr) // functions.go
	if err != nil || host.ID < 1 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": errInvalidHostID.Error()})
		return
	}

	var payload map[string]any
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	rawDeviceType, ok := payload["deviceType"]
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "missing deviceType"})
		return
	}

	deviceType, ok := rawDeviceType.(string)
	if !ok || !models.IsValidDeviceType(deviceType) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid deviceType"})
		return
	}

	oldDeviceType := host.DeviceType
	updatedHost, err := gdb.UpdateDeviceType(host.ID, deviceType)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": errInvalidHostID.Error()})
		return
	}
	if err != nil {
		slog.Error("Failed to update host device type", "id", host.ID, "err", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to update host device type"})
		return
	}

	if oldDeviceType != updatedHost.DeviceType {
		gdb.RecordHostEvent(updatedHost, models.EventDeviceTypeChanged, oldDeviceType, updatedHost.DeviceType)
	}

	c.IndentedJSON(http.StatusOK, updatedHost)
}

// delHost godoc
// @Summary      Delete host
// @Description  Remove a host from the database by its unique ID
// @Tags         hosts
// @Produce      json
// @Param        id   path      string  true  "Host ID"
// @Success      200  {string}  string  "OK"
// @Router       /host/del/{id} [get]
func delHost(c *gin.Context) {
	idStr := c.Param("id")
	host, err := getHostByID(idStr) // functions.go
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gdb.DeleteHostDeviceChangeEvents(host.ID)
	gdb.Delete("now", host.ID)
	slog.Info("Deleting from DB", "host", host)
	c.IndentedJSON(http.StatusOK, "OK")
}

// addHost godoc
// @Summary      Add host manually
// @Description  Add host by MAC, with optional Name, IP, Hardware
// @Description  Returns `models.Host` with this MAC form DB, either just added or existing
// @Tags         hosts
// @Produce      json
// @Param        mac   path      string  true   "Host MAC"
// @Param        name  query     string  false  "Name"
// @Param        ip    query     string  false  "IP"
// @Param        hw    query     string  false  "Hardware"
// @Success      200  {object}  models.Host
// @Router       /host/add/{mac} [get]
func addHost(c *gin.Context) {

	mac := c.Param("mac")
	hosts := gdb.SelectByMAC("now", mac)

	if len(hosts) > 0 {
		slog.Warn("Host with this MAC already exists", "host", hosts[0])
	} else {
		var host models.Host

		host.Mac = mac
		host.Name = c.Query("name")
		host.IP = c.Query("ip")
		host.Hw = c.Query("hw")

		gdb.Update("now", host)
		hosts = gdb.SelectByMAC("now", mac)
		if len(hosts) > 0 {
			gdb.RecordHostEvent(hosts[0], models.EventDiscovered, "", "")
		}

		slog.Info("Added host to DB", "host", hosts[0])
	}

	c.IndentedJSON(http.StatusOK, hosts[0])
}

// editHost godoc
// @Summary      Edit host
// @Description  Update a host's name and optionally toggle its "known" status
// @Tags         hosts
// @Produce      json
// @Param        id     path      string  true  "Host ID"
// @Param        name   path      string  true  "New name for the host"
// @Param        known  path      string  false "Pass 'toggle' to flip the known/unknown status"
// @Success      200    {string}  string  "OK"
// @Router       /edit/{id}/{name}/{known} [get]
func editHost(c *gin.Context) {

	idStr := c.Param("id")
	name := c.Param("name")
	toggleKnown := c.Param("known")

	host, err := getHostByID(idStr) // functions.go
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	host.Name = name
	oldKnown := host.Known

	if toggleKnown == "/toggle" {
		host.Known = 1 - host.Known
	}

	if err := gdb.UpdateWithError("now", host); err != nil {
		slog.Error("Failed to update host", "id", host.ID, "err", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to update host"})
		return
	}

	if oldKnown != host.Known {
		eventType := models.EventUnknown
		if host.Known == 1 {
			eventType = models.EventKnown
		}
		gdb.RecordHostEvent(host, eventType, "", "")
	}

	c.IndentedJSON(http.StatusOK, "OK")
}
