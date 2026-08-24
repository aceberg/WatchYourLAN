package arp

import (
	"context"
	"log/slog"
	"os/exec"
	"strings"
	"time"

	"github.com/aceberg/WatchYourLAN/internal/check"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

var scanCommandTimeout = 2 * time.Minute
var commandRunner = runCommand

func scanIface(iface, scanArgs string) (string, bool) {
	args := []string{"-glNx"}
	args = append(args, strings.Fields(scanArgs)...)
	args = append(args, "-I", iface)

	return commandRunner("arp-scan", args...)
}

func scanStr(str string) (string, bool) {

	args := strings.Fields(str)
	if len(args) == 0 {
		return "", true
	}

	return commandRunner("arp-scan", args...)
}

func runCommand(name string, args ...string) (string, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), scanCommandTimeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, name, args...)

	out, err := cmd.Output()
	slog.Debug(cmd.String())

	if ctx.Err() == context.DeadlineExceeded {
		slog.Error("Command timed out", "cmd", cmd.String(), "timeout", scanCommandTimeout.String())
		return string(""), false
	}

	if check.IfError(err) {
		return string(""), false
	}
	return string(out), true
}

func parseOutput(text, iface string) []models.Host {
	var foundHosts = []models.Host{}

	p := strings.Split(text, "\n")

	for _, host := range p {
		host = strings.TrimSpace(host)
		if host == "" {
			continue
		}

		p := strings.Split(host, "	")
		if len(p) < 3 {
			slog.Warn("Ignoring malformed arp-scan row", "iface", iface, "row", host)
			continue
		}

		var oneHost models.Host
		oneHost.Iface = iface
		oneHost.IP = strings.TrimSpace(p[0])
		oneHost.Mac = strings.TrimSpace(p[1])
		oneHost.Hw = strings.TrimSpace(strings.Join(p[2:], "	"))
		if oneHost.IP == "" || oneHost.Mac == "" || oneHost.Hw == "" {
			slog.Warn("Ignoring incomplete arp-scan row", "iface", iface, "row", host)
			continue
		}
		oneHost.Date = time.Now().Format("2006-01-02 15:04:05")
		oneHost.Now = 1
		foundHosts = append(foundHosts, oneHost)
	}

	return foundHosts
}

// Scan all interfaces
func Scan(ifaces, args string, strs []string) ([]models.Host, bool) {
	var text string
	var p []string
	var foundHosts = []models.Host{}
	scanOK := true

	if ifaces != "" {

		p = strings.Fields(ifaces)

		for _, iface := range p {
			slog.Debug("Scanning interface " + iface)
			var ok bool
			text, ok = scanIface(iface, args)
			if !ok {
				scanOK = false
				continue
			}
			slog.Debug("Found IPs: \n" + text)

			foundHosts = append(foundHosts, parseOutput(text, iface)...)
		}
	}

	for _, s := range strs {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}

		slog.Debug("Scanning string " + s)
		var ok bool
		text, ok = scanStr(s)
		if !ok {
			scanOK = false
			continue
		}
		slog.Debug("Found IPs: \n" + text)
		p = strings.Fields(s)

		foundHosts = append(foundHosts, parseOutput(text, p[len(p)-1])...)
	}

	return foundHosts, scanOK
}
