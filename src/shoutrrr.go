package main

import (
	"log"
	"github.com/containrrr/shoutrrr"
)

func shoutr_notify(message string) {
	if AppConfig.ShoutUrl != "" {
		err := shoutrrr.Send(AppConfig.ShoutUrl, message)
		if err != nil {
			log.Println("ERROR: Notification failed (shoutrrr):", err)
		}
	}
}