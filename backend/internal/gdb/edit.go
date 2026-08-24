package gdb

import (
	"errors"
	"log/slog"

	"github.com/aceberg/WatchYourLAN/internal/check"
	"github.com/aceberg/WatchYourLAN/internal/models"

	"gorm.io/gorm"
)

// Update - update or create host
func Update(table string, oneHost models.Host) {

	check.IfError(UpdateWithError(table, oneHost))
}

// UpdateWithError updates or creates a host and returns persistence errors.
func UpdateWithError(table string, oneHost models.Host) error {

	activeDB, release, err := acquireDB()
	if err != nil {
		return err
	}
	defer release()

	tab := activeDB.Table(table)
	result := tab.Save(&oneHost)

	return result.Error
}

// UpdateDeviceType updates only the manual DeviceType field for a host.
func UpdateDeviceType(id int, deviceType string) (models.Host, error) {
	var host models.Host

	activeDB, release, err := acquireDB()
	if err != nil {
		return host, err
	}
	defer release()

	tab := activeDB.Table("now")
	result := tab.Model(&models.Host{}).
		Where("\"ID\" = ?", id).
		Update("DEVICE_TYPE", deviceType)
	if result.Error != nil {
		return host, result.Error
	}
	if result.RowsAffected == 0 {
		return host, gorm.ErrRecordNotFound
	}

	err = tab.First(&host, id).Error
	return host, err
}

// Delete - delete host from DB
func Delete(table string, id int) {

	activeDB, release, err := acquireDB()
	if err != nil {
		check.IfError(err)
		return
	}
	defer release()

	tab := activeDB.Table(table)
	result := tab.Delete(&models.Host{}, id)
	check.IfError(result.Error)
}

// DeleteOldHistory - delete a list of hosts from History
func DeleteOldHistory(date string) int64 {

	activeDB, release, err := acquireDB()
	if err != nil {
		check.IfError(err)
		return 0
	}
	defer release()

	tab := activeDB.Table("history")
	result := tab.Where("\"DATE\" < ?", date).Delete(&models.Host{})
	check.IfError(result.Error)

	return result.RowsAffected
}

// AddEvent stores a validated host activity event.
func AddEvent(event models.HostEvent) error {

	if !models.IsValidHostEventType(event.EventType) {
		return errors.New("invalid host event type")
	}

	activeDB, release, err := acquireDB()
	if err != nil {
		return err
	}
	defer release()

	tab := activeDB.Table("events")
	result := tab.Create(&event)

	return result.Error
}

// RecordHostEvent stores an activity event and logs failures without failing callers.
func RecordHostEvent(host models.Host, eventType models.HostEventType, oldValue, newValue string) {

	err := AddEvent(models.NewHostEvent(host, eventType, oldValue, newValue))
	if err != nil {
		slog.Error("Failed to record host event", "eventType", eventType, "mac", host.Mac, "err", err)
	}
}

// DeleteOldConnectivityEvents removes online/offline activity events older than date.
func DeleteOldConnectivityEvents(date string) int64 {

	activeDB, release, err := acquireDB()
	if err != nil {
		check.IfError(err)
		return 0
	}
	defer release()

	tab := activeDB.Table("events")
	result := tab.
		Where("\"DATE\" < ?", date).
		Where("\"EVENT_TYPE\" IN ?", []string{
			string(models.EventOnline),
			string(models.EventOffline),
		}).
		Delete(&models.HostEvent{})
	check.IfError(result.Error)

	return result.RowsAffected
}

// DeleteHostDeviceChangeEvents removes persistent device-change events for a deleted host record.
func DeleteHostDeviceChangeEvents(hostID int) int64 {

	activeDB, release, err := acquireDB()
	if err != nil {
		check.IfError(err)
		return 0
	}
	defer release()

	tab := activeDB.Table("events")
	result := tab.
		Where("\"HOST_ID\" = ?", hostID).
		Where("\"EVENT_TYPE\" IN ?", []string{
			string(models.EventDiscovered),
			string(models.EventKnown),
			string(models.EventUnknown),
			string(models.EventDeviceTypeChanged),
		}).
		Delete(&models.HostEvent{})
	check.IfError(result.Error)

	return result.RowsAffected
}

// Clear - delete all hosts from table
func Clear(table string) {

	activeDB, release, err := acquireDB()
	if err != nil {
		check.IfError(err)
		return
	}
	defer release()

	tab := activeDB.Table(table)
	result := tab.Where("1 = 1").Delete(&models.Host{})
	check.IfError(result.Error)
}
