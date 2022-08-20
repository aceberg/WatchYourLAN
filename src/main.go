package main

import (
    "fmt"
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
}

var AppConfig Conf

func main() {
    AppConfig = get_config("./data/config")

    foundHosts := parse_output(scan_iface(AppConfig.Iface))

    //fmt.Println("Found hosts:", foundHosts)

    db_create(AppConfig.DbPath)
    dbHosts := db_select(AppConfig.DbPath)

    //fmt.Println("DB hosts:", dbHosts)

    db_compare(foundHosts, dbHosts)

    dbHosts = db_select(AppConfig.DbPath)
    allHosts := append(foundHosts,dbHosts...)

    fmt.Println(fmt.Sprintf("http://%s:%s", AppConfig.GuiIP, AppConfig.GuiPort))
    
    webgui(AppConfig, allHosts)
}