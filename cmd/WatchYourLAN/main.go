package main

import (
	"fmt"
	// "time"
	"github.com/aceberg/WatchYourLAN/pkg/conf"
	. "github.com/aceberg/WatchYourLAN/pkg/models"
)

var AppConfig Conf

// var AllHosts []Host

// func scan_and_compare() {
//     var foundHosts []Host
//     var dbHosts []Host
//     for {                                  // Endless
//         foundHosts = arp_scan()            // Scan interfaces
//         dbHosts = db_select()              // Select everything from DB
//         db_setnow()                        // Mark hosts in DB as offline
//         hosts_compare(foundHosts, dbHosts) // Compare hosts online and in DB
//                                            // and add them to DB
//         AllHosts = db_select()
//         time.Sleep(time.Duration(AppConfig.Timeout) * time.Second) // Timeout
//     }
// }

func main() {
	// AllHosts = []Host{}
	AppConfig = conf.GetConfig() // Get config from Defaults, Config file, Env

	fmt.Println("CONF =", AppConfig)

	// db_create() // Check if DB exists. Create if not

	// go scan_and_compare()

	// webgui() // Start web GUI
}
