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
}

type Conf struct {
    Iface   string
    DbPath  string
    GuiIP   string
    GuiPort string
}

func main() {
	// iface := "virbr-bw"
    // db_path := "data/hosts.db"

    newConfig := get_config("./data/config")

    text := scan_iface(newConfig.Iface)
    foundHosts := parse_output(text)

    fmt.Println(foundHosts)

    db_create(newConfig.DbPath)

    fmt.Println("http://localhost:8840")
    webgui(foundHosts)
}