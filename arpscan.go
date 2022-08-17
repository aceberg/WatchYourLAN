package main

import (
    // "fmt"
    "log"
    "os/exec"
    "strings"
	"time"
)

func scan_iface(iface string) (string) {
    cmd, err := exec.Command("arp-scan", "-glNx", "-I", iface).Output()
    if err != nil {
        log.Fatal(err)
    }

    return string(cmd)
    // fmt.Println(text)
}

func parse_output(text string) ([]Host) {    
    var foundHosts = []Host{}

    perString := strings.Split(text, "\n")
	currentTime := time.Now()

    for _, v := range perString {
        if v != "" {
            var oneHost Host
            p := strings.Split(v, "	")
            oneHost.Ip = p[0]
            oneHost.Mac = p[1]
            oneHost.Hw = p[2]
			oneHost.Date = currentTime.Format("2006-01-02 15:04:05")
            foundHosts = append(foundHosts, oneHost)
        }
    }

    return foundHosts
    // fmt.Println(foundHosts)
}