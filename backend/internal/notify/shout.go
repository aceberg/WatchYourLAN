package notify

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/nicholas-fedor/shoutrrr"

	"github.com/aceberg/WatchYourLAN/internal/conf"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

// Unknown - send message to log and shoutrrr
func Unknown(host models.Host) {

	msg := fmt.Sprintf("Unknown host found. Name: '%s', IP: '%s', MAC: '%s', Hw: '%s', Iface: '%s'", host.DNS, host.IP, host.Mac, host.Hw, host.Iface)

	slog.Warn(msg)
	shout(msg)
}

// Test Shoutrrr notification
func Test() {

	msg := "test notification"
	slog.Info("Sending " + msg)
	shout(msg)
}

// shout - send msg to Shoutrrr
func shout(msg string) {

	hostname, _ := os.Hostname()
	wyl := "WatchYourLAN on '" + hostname + "': "

	if conf.AppConfig.ShoutURL != "" {
		err := shoutrrr.Send(conf.AppConfig.ShoutURL, wyl+msg)
		if err != nil {
			slog.Error("Notification failed (shoutrrr): ", "", err)
		}
	}
}
