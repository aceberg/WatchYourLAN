package models

import "time"

// HostEventType is a validated machine-readable activity event type.
type HostEventType string

const (
	EventDiscovered        HostEventType = "discovered"
	EventOnline            HostEventType = "online"
	EventOffline           HostEventType = "offline"
	EventKnown             HostEventType = "known"
	EventUnknown           HostEventType = "unknown"
	EventDeviceTypeChanged HostEventType = "device-type-changed"
)

// HostEventTypeValues lists every activity event type accepted for persistence.
var HostEventTypeValues = []HostEventType{
	EventDiscovered,
	EventOnline,
	EventOffline,
	EventKnown,
	EventUnknown,
	EventDeviceTypeChanged,
}

var validHostEventTypes = func() map[string]struct{} {
	values := make(map[string]struct{}, len(HostEventTypeValues))
	for _, value := range HostEventTypeValues {
		values[string(value)] = struct{}{}
	}
	return values
}()

// IsValidHostEventType reports whether value is a supported HostEventType.
func IsValidHostEventType(value string) bool {
	_, ok := validHostEventTypes[value]
	return ok
}

// HostEvent is a persistent snapshot of a meaningful host activity event.
type HostEvent struct {
	ID         int    `gorm:"column:ID;primaryKey"`
	HostID     int    `gorm:"column:HOST_ID"`
	Mac        string `gorm:"column:MAC"`
	Name       string `gorm:"column:NAME"`
	EventType  string `gorm:"column:EVENT_TYPE"`
	Date       string `gorm:"column:DATE"`
	IP         string `gorm:"column:IP"`
	Iface      string `gorm:"column:IFACE"`
	DeviceType string `gorm:"column:DEVICE_TYPE"`
	OldValue   string `gorm:"column:OLD_VALUE"`
	NewValue   string `gorm:"column:NEW_VALUE"`
}

// ActivityStats summarizes retained host activity events.
type ActivityStats struct {
	Total             int64
	Online            int64
	Offline           int64
	Discovered        int64
	Known             int64
	Unknown           int64
	DeviceTypeChanged int64
}

// ActivityDeviceOption identifies a device represented in retained activity events.
type ActivityDeviceOption struct {
	HostID     int
	Mac        string
	Name       string
	DeviceType string
	Exists     bool
}

// NewHostEvent creates a snapshot event from the current Host fields.
func NewHostEvent(host Host, eventType HostEventType, oldValue, newValue string) HostEvent {
	return HostEvent{
		HostID:     host.ID,
		Mac:        host.Mac,
		Name:       host.Name,
		EventType:  string(eventType),
		Date:       time.Now().Format("2006-01-02 15:04:05"),
		IP:         host.IP,
		Iface:      host.Iface,
		DeviceType: host.DeviceType,
		OldValue:   oldValue,
		NewValue:   newValue,
	}
}
