package arp

import (
	"log/slog"
	"os/exec"
	"strings"
	"time"

	"github.com/aceberg/WatchYourLAN/internal/check"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

var arpArgs string

func scanIface(iface string) string {
	var cmd *exec.Cmd

	if arpArgs != "" {
		cmd = exec.Command("arp-scan", "-glNx", arpArgs, "-I", iface)
	} else {
		cmd = exec.Command("arp-scan", "-glNx", "-I", iface)
	}
	out, err := cmd.Output()
	slog.Debug(cmd.String())

	if check.IfError(err) {
		return string("")
	}
	return string(out)
}

func parseOutput(text, iface string) []models.Host {
	var foundHosts = []models.Host{}

	perString := strings.Split(text, "\n")

	for _, host := range perString {
		if host != "" {
			var oneHost models.Host
			p := strings.Split(host, "	")
			oneHost.Iface = iface
			oneHost.IP = p[0]
			oneHost.Mac = p[1]
			oneHost.Hw = p[2]
			oneHost.Date = time.Now().Format("2006-01-02 15:04:05")
			oneHost.Now = 1
			foundHosts = append(foundHosts, oneHost)
		}
	}

	return foundHosts
}

// Scan all interfaces
func Scan(ifaces, args string) []models.Host {
	var text string
	var foundHosts = []models.Host{}
	arpArgs = args

	perString := strings.Split(ifaces, " ")

	for _, iface := range perString {
		slog.Debug("Scanning interface " + iface)
		text = scanIface(iface)
		slog.Debug("Found IPs: \n" + text)

		foundHosts = append(foundHosts, parseOutput(text, iface)...)
	}

	return foundHosts
}
