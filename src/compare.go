package main

// import (
// 	"fmt"
// )

func host_in_db(host Host, dbHosts []Host) bool {
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
	// fmt.Println("Found hosts:", foundHosts)
	// fmt.Println("DB hosts:", dbHosts)

	for _, oneHost := range foundHosts {
		if !(host_in_db(oneHost, dbHosts)) {
			//fmt.Println("No such host!")
			oneHost.Now = 1
			db_insert(oneHost)
		}
	}
}