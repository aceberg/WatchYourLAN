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

func scan() {
    for {
        foundHosts := parse_output(scan_iface(AppConfig.Iface))
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
    
    // fmt.Println("Timeout: ", AppConfig.Timeout)
    go scan()

    webgui()
}