package check

import (
	"bytes"
	"context"
	"encoding/xml"
	"io"
	"log/slog"
	"net"
	"net/http"
	"net/url"
	"os/exec"
	"strings"
	"time"

	"github.com/aceberg/WatchYourLAN/internal/models"
)

// DNS - returns DNS names of a host
func DNS(host models.Host) (name, dns string) {
	names := lookupHostNames(host.IP)

	if len(names) > 0 {
		name = names[0]
		dns = strings.Join(names, " ")
	}

	return name, dns
}

type hostIdentity struct {
	names    []string
	hardware string
}

// EnrichHosts fills host names from local DNS, mDNS and SSDP/UPnP discovery.
func EnrichHosts(hosts []models.Host) []models.Host {
	ipSet := make(map[string]struct{})

	for _, host := range hosts {
		if net.ParseIP(host.IP) != nil {
			ipSet[host.IP] = struct{}{}
		}
	}

	avahi := discoverAvahiBrowse(ipSet)
	ssdp := discoverSSDP(ipSet)

	for i := range hosts {
		names := lookupHostNames(hosts[i].IP)

		if identity, ok := avahi[hosts[i].IP]; ok {
			names = appendUnique(names, identity.names...)
		}

		if identity, ok := ssdp[hosts[i].IP]; ok {
			names = appendUnique(names, identity.names...)

			if isUnknownHardware(hosts[i].Hw) && identity.hardware != "" {
				hosts[i].Hw = identity.hardware
			}
		}

		if len(names) > 0 {
			hosts[i].Name = names[0]
			hosts[i].DNS = strings.Join(names, " ")
		}
	}

	return hosts
}

func lookupHostNames(ip string) []string {
	names := []string{}

	dnsNames, _ := net.LookupAddr(ip)
	names = appendUnique(names, dnsNames...)
	names = appendUnique(names, getentNames(ip)...)
	names = appendUnique(names, avahiNames(ip)...)

	return names
}

func getentNames(ip string) []string {
	out, ok := runOutput(700*time.Millisecond, "getent", "hosts", ip)
	if !ok {
		return nil
	}

	fields := strings.Fields(out)
	if len(fields) < 2 {
		return nil
	}

	return fields[1:]
}

func avahiNames(ip string) []string {
	out, ok := runOutput(900*time.Millisecond, "avahi-resolve-address", ip)
	if !ok {
		return nil
	}

	fields := strings.Fields(out)
	if len(fields) < 2 {
		return nil
	}

	return fields[1:]
}

func runOutput(timeout time.Duration, name string, args ...string) (string, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	out, err := exec.CommandContext(ctx, name, args...).Output()
	if err != nil {
		return "", false
	}

	return string(out), true
}

func discoverAvahiBrowse(targetIPs map[string]struct{}) map[string]hostIdentity {
	identities := make(map[string]hostIdentity)
	if len(targetIPs) == 0 {
		return identities
	}

	out, ok := runOutput(8*time.Second, "avahi-browse", "-artp")
	if !ok {
		return identities
	}

	for _, line := range strings.Split(out, "\n") {
		if !strings.HasPrefix(line, "=") {
			continue
		}

		parts := strings.Split(line, ";")
		if len(parts) < 9 {
			continue
		}

		ip := strings.TrimSpace(parts[7])
		if _, ok := targetIPs[ip]; !ok {
			continue
		}

		identity := identities[ip]
		identity.names = appendUnique(identity.names, parts[3], parts[6])
		identities[ip] = identity
	}

	return identities
}

func discoverSSDP(targetIPs map[string]struct{}) map[string]hostIdentity {
	identities := make(map[string]hostIdentity)
	if len(targetIPs) == 0 {
		return identities
	}

	addr, err := net.ResolveUDPAddr("udp4", "239.255.255.250:1900")
	if err != nil {
		return identities
	}

	conn, err := net.ListenPacket("udp4", ":0")
	if err != nil {
		slog.Debug("SSDP listen failed", "err", err)
		return identities
	}
	defer conn.Close()

	message := strings.Join([]string{
		"M-SEARCH * HTTP/1.1",
		"HOST:239.255.255.250:1900",
		`MAN:"ssdp:discover"`,
		"MX:1",
		"ST:ssdp:all",
		"",
		"",
	}, "\r\n")

	for i := 0; i < 2; i++ {
		_, _ = conn.WriteTo([]byte(message), addr)
	}

	_ = conn.SetDeadline(time.Now().Add(3 * time.Second))
	seenLocations := make(map[string]struct{})

	for {
		buffer := make([]byte, 65535)
		n, remote, err := conn.ReadFrom(buffer)
		if err != nil {
			break
		}

		responseIP, _, err := net.SplitHostPort(remote.String())
		if err != nil {
			responseIP = remote.String()
		}

		headers := parseSSDPHeaders(buffer[:n])
		location := strings.TrimSpace(headers["location"])
		if location == "" {
			continue
		}
		if _, seen := seenLocations[location]; seen {
			continue
		}
		seenLocations[location] = struct{}{}

		ip := matchingLocationIP(location, responseIP, targetIPs)
		if ip == "" {
			continue
		}

		identity := identities[ip]
		desc := fetchSSDPDescription(location)
		identity.names = appendUnique(identity.names, desc["friendlyName"])
		identity.hardware = joinNonEmpty(desc["manufacturer"], desc["modelName"], desc["modelNumber"])
		identities[ip] = identity
	}

	return identities
}

func parseSSDPHeaders(payload []byte) map[string]string {
	headers := make(map[string]string)

	for _, line := range strings.Split(string(payload), "\n") {
		key, value, ok := strings.Cut(line, ":")
		if !ok {
			continue
		}
		headers[strings.ToLower(strings.TrimSpace(key))] = strings.TrimSpace(value)
	}

	return headers
}

func matchingLocationIP(location, responseIP string, targetIPs map[string]struct{}) string {
	if _, ok := targetIPs[responseIP]; ok {
		return responseIP
	}

	parsed, err := url.Parse(location)
	if err != nil || parsed.Hostname() == "" {
		return ""
	}

	locationIP := net.ParseIP(parsed.Hostname())
	if locationIP != nil {
		ip := locationIP.String()
		if _, ok := targetIPs[ip]; ok {
			return ip
		}
		return ""
	}

	ips, err := net.LookupHost(parsed.Hostname())
	if err != nil {
		return ""
	}
	for _, ip := range ips {
		if _, ok := targetIPs[ip]; ok {
			return ip
		}
	}

	return ""
}

func fetchSSDPDescription(location string) map[string]string {
	result := make(map[string]string)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, location, nil)
	if err != nil {
		return result
	}
	req.Header.Set("User-Agent", "WatchYourLAN/2.1.4")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return result
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(io.LimitReader(resp.Body, 256*1024))
	if err != nil {
		return result
	}

	return parseSSDPDescription(body)
}

func parseSSDPDescription(body []byte) map[string]string {
	result := make(map[string]string)
	wanted := map[string]struct{}{
		"friendlyName": {},
		"manufacturer": {},
		"modelName":    {},
		"modelNumber":  {},
	}

	decoder := xml.NewDecoder(bytes.NewReader(body))
	current := ""

	for {
		token, err := decoder.Token()
		if err != nil {
			break
		}

		switch item := token.(type) {
		case xml.StartElement:
			if _, ok := wanted[item.Name.Local]; ok {
				current = item.Name.Local
			}
		case xml.CharData:
			if current != "" {
				result[current] = cleanName(string(item))
			}
		case xml.EndElement:
			if item.Name.Local == current {
				current = ""
			}
		}
	}

	return result
}

func appendUnique(items []string, more ...string) []string {
	seen := make(map[string]struct{}, len(items)+len(more))
	out := make([]string, 0, len(items)+len(more))

	for _, item := range append(items, more...) {
		item = cleanName(item)
		if item == "" {
			continue
		}
		if _, ok := seen[item]; ok {
			continue
		}
		seen[item] = struct{}{}
		out = append(out, item)
	}

	return out
}

func cleanName(value string) string {
	value = strings.TrimSpace(value)
	value = strings.TrimSuffix(value, ".")
	value = strings.Join(strings.Fields(value), " ")

	if value == "" || net.ParseIP(value) != nil {
		return ""
	}

	return value
}

func isUnknownHardware(value string) bool {
	value = strings.TrimSpace(value)
	return value == "" || strings.HasPrefix(value, "(Unknown")
}

func joinNonEmpty(parts ...string) string {
	out := []string{}
	for _, part := range parts {
		part = cleanName(part)
		if part != "" {
			out = append(out, part)
		}
	}

	return strings.Join(out, " ")
}
