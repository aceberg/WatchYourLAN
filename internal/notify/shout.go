package notify

import (
	"github.com/containrrr/shoutrrr"
	"log/slog"
)

// Shout - send message with shoutrrr
func Shout(message string, url string) {
	if url != "" {
		err := shoutrrr.Send(url, message)
		if err != nil {
			slog.Error("Notification failed (shoutrrr): ", err)
		}
	}
}
