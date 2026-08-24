package gdb

import (
	"github.com/aceberg/WatchYourLAN/internal/check"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

// EventQuery describes filters for persistent activity event selection.
type EventQuery struct {
	Limit      int
	Offset     int
	Mac        string
	Macs       []string
	HostID     int
	EventTypes []models.HostEventType
}

// Select - get all hosts
func Select(table string) (dbHosts []models.Host, ok bool) {

	activeDB, release, err := acquireDB()
	if err != nil {
		return dbHosts, !check.IfError(err)
	}
	defer release()

	tab := activeDB.Table(table)
	err = tab.Find(&dbHosts).Error

	return dbHosts, !check.IfError(err)
}

// SelectByID - get host by ID
func SelectByID(id int) (host models.Host) {

	activeDB, release, err := acquireDB()
	if err != nil {
		check.IfError(err)
		return host
	}
	defer release()

	tab := activeDB.Table("now")
	tab.First(&host, id)

	return host
}

// SelectByMAC - get all hosts by MAC
func SelectByMAC(table, mac string) (hosts []models.Host) {

	activeDB, release, err := acquireDB()
	if err != nil {
		check.IfError(err)
		return hosts
	}
	defer release()

	tab := activeDB.Table(table)
	tab.Where("\"MAC\" = ?", mac).Find(&hosts)

	return hosts
}

// SelectByDate - get all hosts by MAC and DATE
func SelectByDate(mac, date string) (hosts []models.Host) {

	activeDB, release, err := acquireDB()
	if err != nil {
		check.IfError(err)
		return hosts
	}
	defer release()

	tab := activeDB.Table("history")
	tab.
		Where("\"MAC\" = ?", mac).
		Where("\"DATE\" LIKE ?", date+"%").
		Find(&hosts)

	return hosts
}

// SelectLatest - get latest hosts by MAC
func SelectLatest(mac string, number int) (hosts []models.Host) {

	activeDB, release, err := acquireDB()
	if err != nil {
		check.IfError(err)
		return hosts
	}
	defer release()

	tab := activeDB.Table("history")
	tab.
		Where("\"MAC\" = ?", mac).
		Order("\"DATE\" DESC").
		Limit(number).
		Find(&hosts)

	return hosts
}

// SelectEvents returns recent host activity events, newest first.
func SelectEvents(limit int, mac string) (events []models.HostEvent, ok bool) {

	return SelectEventsFiltered(EventQuery{
		Limit: limit,
		Mac:   mac,
	})
}

// SelectEventsFiltered returns recent host activity events matching query, newest first.
func SelectEventsFiltered(query EventQuery) (events []models.HostEvent, ok bool) {

	activeDB, release, err := acquireDB()
	if err != nil {
		return events, !check.IfError(err)
	}
	defer release()

	tab := activeDB.Table("events")
	macs := eventQueryMacs(query)
	if len(macs) > 0 {
		tab = tab.Where("\"MAC\" IN ?", macs)
	}
	if query.HostID > 0 {
		tab = tab.Where("\"HOST_ID\" = ?", query.HostID)
	}
	if len(query.EventTypes) > 0 {
		eventTypes := make([]string, 0, len(query.EventTypes))
		for _, eventType := range query.EventTypes {
			eventTypes = append(eventTypes, string(eventType))
		}
		tab = tab.Where("\"EVENT_TYPE\" IN ?", eventTypes)
	}
	err = tab.
		Order("\"DATE\" DESC").
		Order("\"ID\" DESC").
		Limit(query.Limit).
		Offset(query.Offset).
		Find(&events).Error

	return events, !check.IfError(err)
}

// SelectEventStats returns faceted counts for retained activity events.
func SelectEventStats(macs []string) (stats models.ActivityStats, ok bool) {

	type eventStatRow struct {
		EventType string `gorm:"column:EVENT_TYPE"`
		Count     int64  `gorm:"column:count"`
	}

	activeDB, release, err := acquireDB()
	if err != nil {
		return stats, !check.IfError(err)
	}
	defer release()

	tab := activeDB.Table("events")
	macs = normalizeMacs(macs)
	if len(macs) > 0 {
		tab = tab.Where("\"MAC\" IN ?", macs)
	}

	var rows []eventStatRow
	err = tab.
		Select("\"EVENT_TYPE\", COUNT(*) as count").
		Group("EVENT_TYPE").
		Scan(&rows).Error
	if check.IfError(err) {
		return stats, false
	}

	for _, row := range rows {
		stats.Total += row.Count
		switch models.HostEventType(row.EventType) {
		case models.EventOnline:
			stats.Online = row.Count
		case models.EventOffline:
			stats.Offline = row.Count
		case models.EventDiscovered:
			stats.Discovered = row.Count
		case models.EventKnown:
			stats.Known = row.Count
		case models.EventUnknown:
			stats.Unknown = row.Count
		case models.EventDeviceTypeChanged:
			stats.DeviceTypeChanged = row.Count
		}
	}

	return stats, true
}

// SelectActivityDeviceOptions returns current hosts and retained event-only devices.
func SelectActivityDeviceOptions() (devices []models.ActivityDeviceOption, ok bool) {

	var hosts []models.Host
	activeDB, release, err := acquireDB()
	if err != nil {
		return devices, !check.IfError(err)
	}
	defer release()

	err = activeDB.Table("now").
		Order("\"NAME\" ASC").
		Order("\"MAC\" ASC").
		Find(&hosts).Error
	if check.IfError(err) {
		return devices, false
	}

	seen := make(map[string]struct{}, len(hosts))
	for _, host := range hosts {
		if host.Mac == "" {
			continue
		}
		seen[host.Mac] = struct{}{}
		devices = append(devices, models.ActivityDeviceOption{
			HostID:     host.ID,
			Mac:        host.Mac,
			Name:       host.Name,
			DeviceType: host.DeviceType,
			Exists:     true,
		})
	}

	var events []models.HostEvent
	err = activeDB.Table("events").
		Order("\"DATE\" DESC").
		Order("\"ID\" DESC").
		Find(&events).Error
	if check.IfError(err) {
		return devices, false
	}

	for _, event := range events {
		if event.Mac == "" {
			continue
		}
		if _, ok := seen[event.Mac]; ok {
			continue
		}
		seen[event.Mac] = struct{}{}
		devices = append(devices, models.ActivityDeviceOption{
			HostID:     event.HostID,
			Mac:        event.Mac,
			Name:       event.Name,
			DeviceType: event.DeviceType,
			Exists:     false,
		})
	}

	return devices, true
}

// SelectEventsByHostID returns recent host activity events for one host, newest first.
func SelectEventsByHostID(hostID int, limit int) (events []models.HostEvent, ok bool) {

	return SelectEventsFiltered(EventQuery{
		Limit:  limit,
		HostID: hostID,
	})
}

func eventQueryMacs(query EventQuery) []string {
	macs := query.Macs
	if query.Mac != "" {
		macs = append(macs, query.Mac)
	}

	return normalizeMacs(macs)
}

func normalizeMacs(macs []string) []string {
	seen := make(map[string]struct{}, len(macs))
	normalized := make([]string, 0, len(macs))
	for _, mac := range macs {
		if mac == "" {
			continue
		}
		if _, ok := seen[mac]; ok {
			continue
		}
		seen[mac] = struct{}{}
		normalized = append(normalized, mac)
	}

	return normalized
}
