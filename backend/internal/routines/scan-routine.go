package routines

import (
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/aceberg/WatchYourLAN/internal/arp"
	"github.com/aceberg/WatchYourLAN/internal/check"
	"github.com/aceberg/WatchYourLAN/internal/conf"
	"github.com/aceberg/WatchYourLAN/internal/gdb"
	"github.com/aceberg/WatchYourLAN/internal/influx"
	"github.com/aceberg/WatchYourLAN/internal/models"
	"github.com/aceberg/WatchYourLAN/internal/notify"
	"github.com/aceberg/WatchYourLAN/internal/prometheus"
)

var scanMu sync.Mutex

// ScanNow runs one scan immediately and waits until DB state is updated.
func ScanNow() {
	scanMu.Lock()
	defer scanMu.Unlock()

	foundHosts := arp.Scan(conf.AppConfig.Ifaces, conf.AppConfig.ArpArgs, conf.AppConfig.ArpStrs)
	foundHosts = check.EnrichHosts(foundHosts)

	compareHosts(newFoundHostIndex(foundHosts))
}

func startScan(quit chan bool) {
	var lastDate, nowDate, plusDate time.Time

	for {
		select {
		case <-quit:
			return
		default:
			nowDate = time.Now()
			plusDate = lastDate.Add(time.Duration(conf.AppConfig.Timeout) * time.Second)

			if nowDate.After(plusDate) {

				ScanNow()

				lastDate = time.Now()
			}

			select {
			case <-quit:
				return
			case <-time.After(time.Duration(1) * time.Minute):
			}
		}
	}
}

func compareHosts(foundHosts *foundHostIndex) {

	allHosts, ok := gdb.Select("now")
	if !ok {
		return
	}

	existingMACCounts := make(map[string]int)
	for _, aHost := range allHosts {
		mac := normalizeMAC(aHost.Mac)
		if mac != "" {
			existingMACCounts[mac]++
		}
	}

	for _, aHost := range allHosts {

		mac := normalizeMAC(aHost.Mac)
		fHost, exists := foundHosts.match(aHost, existingMACCounts[mac] == 1)
		if exists {

			aHost.Iface = fHost.Iface
			aHost.IP = fHost.IP
			if aHost.Name == "" && fHost.Name != "" {
				aHost.Name = fHost.Name
			}
			if aHost.DNS == "" && fHost.DNS != "" {
				aHost.DNS = fHost.DNS
			}
			if isUnknownHardware(aHost.Hw) && fHost.Hw != "" {
				aHost.Hw = fHost.Hw
			}
			aHost.Date = fHost.Date
			aHost.Now = 1

			foundHosts.delete(fHost)

		} else {
			aHost.Now = 0
		}
		gdb.Update("now", aHost)

		aHost.ID = 0
		aHost.Date = time.Now().Format("2006-01-02 15:04:05")
		gdb.Update("history", aHost)

		if conf.AppConfig.InfluxEnable {
			influx.Add(conf.AppConfig, aHost)
		}
		if conf.AppConfig.PrometheusEnable {
			prometheus.Add(aHost)
		}
	}

	for _, fHost := range foundHosts.remaining() {

		if fHost.Name == "" || fHost.DNS == "" {
			fHost.Name, fHost.DNS = check.DNS(fHost)
		}
		notify.Unknown(fHost) // Log and Shoutrrr

		gdb.Update("now", fHost)
	}
}

type foundHostIndex struct {
	hosts     map[string]models.Host
	keysByIP  map[string][]string
	keysByMAC map[string][]string
}

func newFoundHostIndex(hosts []models.Host) *foundHostIndex {
	index := &foundHostIndex{
		hosts:     make(map[string]models.Host),
		keysByIP:  make(map[string][]string),
		keysByMAC: make(map[string][]string),
	}

	for _, host := range hosts {
		key := hostIdentityKey(host)
		if key == "" {
			continue
		}

		index.hosts[key] = host

		if host.IP != "" {
			index.keysByIP[host.IP] = appendUnique(index.keysByIP[host.IP], key)
		}

		mac := normalizeMAC(host.Mac)
		if mac != "" {
			index.keysByMAC[mac] = appendUnique(index.keysByMAC[mac], key)
		}
	}

	return index
}

func (index *foundHostIndex) match(existing models.Host, allowMACFallback bool) (models.Host, bool) {
	key := hostIdentityKey(existing)
	if host, ok := index.hosts[key]; ok {
		return host, true
	}

	if existing.IP != "" {
		if host, ok := index.only(index.keysByIP[existing.IP]); ok {
			return host, true
		}
	}

	mac := normalizeMAC(existing.Mac)
	if allowMACFallback && mac != "" {
		if host, ok := index.only(index.keysByMAC[mac]); ok {
			return host, true
		}
	}

	return models.Host{}, false
}

func (index *foundHostIndex) only(keys []string) (models.Host, bool) {
	active := make([]string, 0, len(keys))
	for _, key := range keys {
		if _, ok := index.hosts[key]; ok {
			active = append(active, key)
		}
	}

	if len(active) != 1 {
		return models.Host{}, false
	}

	host, ok := index.hosts[active[0]]
	return host, ok
}

func (index *foundHostIndex) delete(host models.Host) {
	key := hostIdentityKey(host)
	delete(index.hosts, key)
}

func (index *foundHostIndex) remaining() []models.Host {
	keys := make([]string, 0, len(index.hosts))
	for key := range index.hosts {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return index.hosts[keys[i]].IP < index.hosts[keys[j]].IP
	})

	hosts := make([]models.Host, 0, len(keys))
	for _, key := range keys {
		hosts = append(hosts, index.hosts[key])
	}

	return hosts
}

func hostIdentityKey(host models.Host) string {
	ip := strings.TrimSpace(host.IP)
	mac := normalizeMAC(host.Mac)

	if mac != "" && ip != "" {
		return mac + "|" + ip
	}
	if mac != "" {
		return mac
	}
	return ip
}

func normalizeMAC(value string) string {
	return strings.ToLower(strings.TrimSpace(value))
}

func appendUnique(items []string, more ...string) []string {
	seen := make(map[string]struct{}, len(items)+len(more))
	out := make([]string, 0, len(items)+len(more))

	for _, item := range append(items, more...) {
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

func isUnknownHardware(value string) bool {
	value = strings.TrimSpace(value)
	return value == "" || strings.HasPrefix(value, "(Unknown")
}
