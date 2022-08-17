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

func main() {
    newConfig := get_config("./data/config")

    foundHosts := parse_output(scan_iface(newConfig.Iface))

    fmt.Println("Found hosts:", foundHosts)

    db_create(newConfig.DbPath)
    db_insert(newConfig.DbPath, foundHosts[0])
    db_select(newConfig.DbPath)

    fmt.Println(fmt.Sprintf("http://%s:%s", newConfig.GuiIP, newConfig.GuiPort))
    
    webgui(newConfig, foundHosts)
}