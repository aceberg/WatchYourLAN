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

func main() {
	iface := "virbr-bw"
    db_path := "data/hosts.db"

    text := scan_iface(iface)
    foundHosts := parse_output(text)

    fmt.Println(foundHosts)

    db_create(db_path)

    fmt.Println("http://localhost:8840")
    webgui(foundHosts)
}