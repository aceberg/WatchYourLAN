package main

// import (
//     "fmt"
// )

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
}

var AppConfig Conf
var AllHosts []Host

func main() {
    AppConfig = get_config("./data/config")
    db_create()

    foundHosts := parse_output(scan_iface(AppConfig.Iface))
    dbHosts := db_select()
    db_setnow()
    db_compare(foundHosts, dbHosts)

    AllHosts = db_select()
    
    webgui()
}