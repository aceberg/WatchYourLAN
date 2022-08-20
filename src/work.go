package main

import (
	"fmt"
)

func host_in_db(host Host, dbHosts []Host) bool {
	for _, oneHost := range dbHosts {
		if host.Ip == oneHost.Ip {
			return true
		}
	}
	return false
}

func db_compare(foundHosts []Host, dbHosts []Host) {
	fmt.Println("Found hosts:", foundHosts)
	fmt.Println("DB hosts:", dbHosts)

	for _, oneHost := range foundHosts {
		if host_in_db(oneHost, dbHosts) {
			fmt.Println("Host in db")	
		} else {
			fmt.Println("No such host!")
			oneHost.Now = 0
			db_insert(AppConfig.DbPath, oneHost)
		}
	}
}