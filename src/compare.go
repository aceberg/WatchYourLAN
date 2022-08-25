package main

import (
	"fmt"
	"log"
)

func host_in_db(host Host, dbHosts []Host) bool { // Check if host is already in DB
	for _, oneHost := range dbHosts {
		if host.Ip == oneHost.Ip && host.Mac == oneHost.Mac && host.Hw == oneHost.Hw {
			oneHost.Date = host.Date
			oneHost.Now = 1
			db_update(oneHost)
			return true
		}
	}
	return false
}

func hosts_compare(foundHosts []Host, dbHosts []Host) {
	for _, oneHost := range foundHosts {
		if !(host_in_db(oneHost, dbHosts)) {
			oneHost.Now = 1			// Mark host online
			msg := fmt.Sprintf("UNKNOWN HOST IP: '%s', MAC: '%s', Hw: '%s'", oneHost.Ip, oneHost.Mac, oneHost.Hw)
			log.Println("WARN:", msg)
			shoutr_notify(msg)		// Notify through Shoutrrr
			db_insert(oneHost)
		}
	}
}