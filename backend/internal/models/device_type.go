package models

// DeviceType is the manually assigned classification for a Host.
type DeviceType string

const (
	DeviceTypeUnassigned     DeviceType = ""
	DeviceTypeRouter         DeviceType = "router"
	DeviceTypeSwitch         DeviceType = "switch"
	DeviceTypeAccessPoint    DeviceType = "access-point"
	DeviceTypeFirewall       DeviceType = "firewall"
	DeviceTypeServer         DeviceType = "server"
	DeviceTypeNAS            DeviceType = "nas"
	DeviceTypeDesktop        DeviceType = "desktop"
	DeviceTypeLaptop         DeviceType = "laptop"
	DeviceTypePhone          DeviceType = "phone"
	DeviceTypeTablet         DeviceType = "tablet"
	DeviceTypeTV             DeviceType = "tv"
	DeviceTypePrinter        DeviceType = "printer"
	DeviceTypeCamera         DeviceType = "camera"
	DeviceTypeIoT            DeviceType = "iot"
	DeviceTypeVirtualMachine DeviceType = "virtual-machine"
	DeviceTypeContainer      DeviceType = "container"
	DeviceTypeGameConsole    DeviceType = "game-console"
	DeviceTypeOther          DeviceType = "other"
)

// DeviceTypeValues lists every value accepted by the backend API.
var DeviceTypeValues = []DeviceType{
	DeviceTypeUnassigned,
	DeviceTypeRouter,
	DeviceTypeSwitch,
	DeviceTypeAccessPoint,
	DeviceTypeFirewall,
	DeviceTypeServer,
	DeviceTypeNAS,
	DeviceTypeDesktop,
	DeviceTypeLaptop,
	DeviceTypePhone,
	DeviceTypeTablet,
	DeviceTypeTV,
	DeviceTypePrinter,
	DeviceTypeCamera,
	DeviceTypeIoT,
	DeviceTypeVirtualMachine,
	DeviceTypeContainer,
	DeviceTypeGameConsole,
	DeviceTypeOther,
}

var validDeviceTypes = func() map[string]struct{} {
	values := make(map[string]struct{}, len(DeviceTypeValues))
	for _, value := range DeviceTypeValues {
		values[string(value)] = struct{}{}
	}
	return values
}()

// IsValidDeviceType reports whether value is a supported DeviceType.
func IsValidDeviceType(value string) bool {
	_, ok := validDeviceTypes[value]
	return ok
}
