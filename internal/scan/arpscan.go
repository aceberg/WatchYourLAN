package scan

import (
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/aceberg/WatchYourLAN/internal/models"
)

func scanIface(iface string) string {
	cmd, err := exec.Command("arp-scan", "-glNx", "-I", iface).Output()
	if err != nil {
		return string("")
	}
	return string(cmd)
}

func parseOutput(text string) []models.Host {
	var foundHosts = []models.Host{}

	perString := strings.Split(text, "\n")
	currentTime := time.Now()

	for _, host := range perString {
		if host != "" {
			var oneHost models.Host
			p := strings.Split(host, "	")
			oneHost.IP = p[0]
			oneHost.Mac = p[1]
			oneHost.Hw = p[2]
			oneHost.Date = currentTime.Format("2006-01-02 15:04:05")
			oneHost.Now = 1
			foundHosts = append(foundHosts, oneHost)
		}
	}

	return foundHosts
}

// Scan all interfaces
func arpScan(allIfaces string, logLevel string) []models.Host {
	var text string
	var foundHosts = []models.Host{}

	perString := strings.Split(allIfaces, " ")

	for _, iface := range perString {
		text = scanIface(iface)
		if logLevel != "short" {
			log.Println("INFO: scanning interface", iface)
			log.Println("INFO: found IPs:", text)
		}
		foundHosts = append(foundHosts, parseOutput(text)...)
	}

	return foundHosts
}
