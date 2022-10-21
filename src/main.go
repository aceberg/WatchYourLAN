package main

import (
    "time"
)

type Host struct {
    Id    uint16
    Name  string
    Ip    string
    Mac   string
    Hw    string
    Date  string
    Known uint16
    Now   uint16
}

type Conf struct {
    Iface    string
    DbPath   string
    GuiIP    string
    GuiPort  string
    Timeout  int
    ShoutUrl string
    Theme    string
}

var AppConfig Conf
var AllHosts []Host

func scan_and_compare() {
    var foundHosts []Host
    var dbHosts []Host
    for {                                  // Endless
        foundHosts = arp_scan()            // Scan interfaces
        dbHosts = db_select()              // Select everything from DB
        db_setnow()                        // Mark hosts in DB as offline
        hosts_compare(foundHosts, dbHosts) // Compare hosts online and in DB
                                           // and add them to DB
	update_all_hosts()
        time.Sleep(time.Duration(AppConfig.Timeout) * time.Second) // Timeout
    }
}

func update_all_hosts(){
    AllHosts = db_select() 
}

func main() {
    AllHosts = []Host{}
    AppConfig = get_config() // Get config from Defaults, Config file, Env

    db_create() // Check if DB exists. Create if not
    
    go scan_and_compare()

    webgui() // Start web GUI
}
