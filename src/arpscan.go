package main

import (
	"log"
	"os/exec"
	"strings"
	"sync"
	"time"
)

func scan_iface(iface string) string {
	cmd, err := exec.Command("arp-scan", "-glNx", "-I", iface).Output()
	if err != nil {
		return string("")
	} else {
		return string(cmd)
	}
}

func parse_output(text string) []Host {
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

func arp_scan() []Host {
	var foundHosts = []Host{}
	perString := strings.Split(AppConfig.Iface, " ")
	wg := sync.WaitGroup{}
	wg.Add(len(perString))
	arpScanMutex := &sync.Mutex{}

	for _, iface := range perString {
		// Multi thread arp-scan
		go func(thisInterface string) {
			log.Println("INFO: scanning interface", thisInterface)
			text := scan_iface(thisInterface)
			log.Println("INFO: found IPs:", text)
			parsedHosts := parse_output(text)
			arpScanMutex.Lock()
			foundHosts = append(foundHosts, parsedHosts...)
			arpScanMutex.Unlock()
			wg.Done()
		}(iface)
	}
	wg.Wait()
	return foundHosts
}
