package arp

import (
	"log/slog"
	"net"
	"os"
	"os/exec"
	"sort"
	"strconv"
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

func scanStr(str string) string {

	args := strings.Split(str, " ")
	cmd := exec.Command("arp-scan", args...)

	out, err := cmd.Output()
	slog.Debug(cmd.String())

	if check.IfError(err) {
		return string("")
	}
	return string(out)
}

func parseOutput(text, iface string) []models.Host {
	var foundHosts = []models.Host{}

	p := strings.Split(text, "\n")

	for _, host := range p {
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
func Scan(ifaces, args string, strs []string) []models.Host {
	var text string
	var p []string
	var foundHosts = []models.Host{}
	arpArgs = args

	p = resolveScanInterfaces(ifaces)
	for _, iface := range p {
		slog.Debug("Scanning interface " + iface)
		text = scanIface(iface)
		slog.Debug("Found IPs: \n" + text)

		foundHosts = append(foundHosts, parseOutput(text, iface)...)
	}

	for _, s := range strs {
		slog.Debug("Scanning string " + s)
		text = scanStr(s)
		slog.Debug("Found IPs: \n" + text)
		p = strings.Split(s, " ")

		foundHosts = append(foundHosts, parseOutput(text, p[len(p)-1])...)
	}

	return foundHosts
}

func resolveScanInterfaces(ifaces string) []string {
	configured := strings.Fields(ifaces)
	available := usableInterfaceMap()
	auto := autoScanInterfaces(available)
	selected, invalid, usedAuto := selectScanInterfaces(configured, available, auto)

	if len(invalid) > 0 {
		slog.Warn("Ignoring unavailable scan interfaces", "ifaces", strings.Join(invalid, " "))
	}
	if usedAuto && len(selected) > 0 {
		slog.Warn("Using auto-detected scan interfaces", "ifaces", strings.Join(selected, " "))
	}

	return selected
}

func selectScanInterfaces(configured []string, available map[string]bool, auto []string) ([]string, []string, bool) {
	if len(configured) == 0 {
		return uniqueStrings(auto), nil, len(auto) > 0
	}

	var selected []string
	var invalid []string
	for _, iface := range configured {
		if available[iface] {
			selected = appendUniqueString(selected, iface)
			continue
		}
		invalid = appendUniqueString(invalid, iface)
	}

	if len(selected) > 0 {
		return selected, invalid, false
	}

	return uniqueStrings(auto), invalid, len(auto) > 0
}

func usableInterfaceMap() map[string]bool {
	interfaces, err := net.Interfaces()
	if err != nil {
		slog.Warn("Cannot list network interfaces", "err", err)
		return nil
	}

	available := make(map[string]bool)
	for _, iface := range interfaces {
		if isUsableInterface(iface) {
			available[iface.Name] = true
		}
	}

	return available
}

func isUsableInterface(iface net.Interface) bool {
	if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
		return false
	}

	addrs, err := iface.Addrs()
	if err != nil {
		slog.Debug("Cannot list interface addresses", "iface", iface.Name, "err", err)
		return false
	}

	for _, addr := range addrs {
		if hasUsableIPv4(addr) {
			return true
		}
	}

	return false
}

func hasUsableIPv4(addr net.Addr) bool {
	ipNet, ok := addr.(*net.IPNet)
	if !ok {
		return false
	}

	ip := ipNet.IP.To4()
	return ip != nil && !ip.IsLoopback() && !ip.IsLinkLocalUnicast()
}

func autoScanInterfaces(available map[string]bool) []string {
	if len(available) == 0 {
		return nil
	}

	ifaces := defaultRouteInterfaces(available)
	if len(ifaces) > 0 {
		return ifaces
	}

	for iface := range available {
		if isVirtualInterface(iface) {
			continue
		}
		ifaces = append(ifaces, iface)
	}
	sort.Strings(ifaces)
	return ifaces
}

func defaultRouteInterfaces(available map[string]bool) []string {
	data, err := os.ReadFile("/proc/net/route")
	if err != nil {
		slog.Debug("Cannot read default routes", "err", err)
		return nil
	}

	var ifaces []string
	minMetric := -1
	for _, line := range strings.Split(string(data), "\n")[1:] {
		fields := strings.Fields(line)
		if len(fields) < 7 || fields[1] != "00000000" {
			continue
		}

		iface := fields[0]
		if !available[iface] || isVirtualInterface(iface) {
			continue
		}

		metric, err := strconv.Atoi(fields[6])
		if err != nil {
			metric = 0
		}

		switch {
		case minMetric == -1 || metric < minMetric:
			minMetric = metric
			ifaces = []string{iface}
		case metric == minMetric:
			ifaces = appendUniqueString(ifaces, iface)
		}
	}

	sort.Strings(ifaces)
	return ifaces
}

func isVirtualInterface(iface string) bool {
	prefixes := []string{
		"br-",
		"docker",
		"lxc",
		"tailscale",
		"tap",
		"tun",
		"vboxnet",
		"veth",
		"virbr",
		"vmnet",
		"wg",
	}

	for _, prefix := range prefixes {
		if strings.HasPrefix(iface, prefix) {
			return true
		}
	}

	return false
}

func uniqueStrings(values []string) []string {
	var out []string
	for _, value := range values {
		out = appendUniqueString(out, value)
	}
	return out
}

func appendUniqueString(values []string, value string) []string {
	if value == "" {
		return values
	}

	for _, item := range values {
		if item == value {
			return values
		}
	}

	return append(values, value)
}
