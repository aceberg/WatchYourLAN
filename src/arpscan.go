package main

import (
    // "fmt"
    "os/exec"
    "strings"
	"time"
)

func scan_iface(iface string) (string) {
    cmd, err := exec.Command("arp-scan", "-glNx", "-I", iface).Output()
    if err != nil {
        return string("")
    } else {
        return string(cmd)
    }
}

func parse_output(text string) ([]Host) {    
    var foundHosts = []Host{}

    perString := strings.Split(text, "\n")
	currentTime := time.Now()

    for _, host := range perString {
        if host != "" {
            var oneHost Host
            p := strings.Split(host, "	")
            oneHost.Ip = p[0]
            oneHost.Mac = p[1]
            oneHost.Hw = p[2]
			oneHost.Date = currentTime.Format("2006-01-02 15:04:05")
            oneHost.Now = 1
            foundHosts = append(foundHosts, oneHost)
        }
    }

    return foundHosts
}

func arp_scan() ([]Host) {
    var text string
    var foundHosts = []Host{}

    perString := strings.Split(AppConfig.Iface, " ")

    for _, iface := range perString {
        text = scan_iface(iface)
        foundHosts = append(foundHosts, parse_output(text)...)
    }

    return foundHosts
}