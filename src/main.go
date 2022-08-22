package main

import (
    // "fmt"
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
    Iface   string
    DbPath  string
    GuiIP   string
    GuiPort string
    Timeout int
}

var AppConfig Conf
var AllHosts []Host

func scan_and_compare() {
    for {
        foundHosts := arp_scan()
        dbHosts := db_select()
        db_setnow()
        hosts_compare(foundHosts, dbHosts)

        AllHosts = db_select()
        // fmt.Println("Refresh")
        time.Sleep(time.Duration(AppConfig.Timeout) * time.Second)
    }
}

func main() {
    AllHosts = []Host{}
    AppConfig = get_config("./data/config")
    db_create()
    
    go scan_and_compare()

    webgui()
}