package notify

import (
	"github.com/containrrr/shoutrrr"
	"log"
)

// Shoutrrr - send message with shoutrrr
func Shoutrrr(message string, url string) {
	if url != "" {
		err := shoutrrr.Send(url, message)
		if err != nil {
			log.Println("ERROR: Notification failed (shoutrrr):", err)
		}
	}
}
