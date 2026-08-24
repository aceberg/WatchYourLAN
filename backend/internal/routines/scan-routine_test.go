package routines

import (
	"path/filepath"
	"testing"

	"github.com/aceberg/WatchYourLAN/internal/conf"
	"github.com/aceberg/WatchYourLAN/internal/gdb"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

func setupScanRoutineTest(t *testing.T) {
	t.Helper()

	oldConfig := conf.AppConfig
	conf.AppConfig.UseDB = "sqlite"
	conf.AppConfig.DBPath = filepath.Join(t.TempDir(), "scan-routine-test.db")
	conf.AppConfig.InfluxEnable = false
	conf.AppConfig.PrometheusEnable = false
	conf.AppConfig.ShoutURL = ""
	gdb.Start()

	t.Cleanup(func() {
		if err := gdb.Close(); err != nil {
			t.Errorf("gdb.Close: %v", err)
		}
		conf.AppConfig = oldConfig
	})
}

func TestCompareHostsPreservesDeviceType(t *testing.T) {
	setupScanRoutineTest(t)

	gdb.Update("now", models.Host{
		Name:       "router",
		Iface:      "eth0",
		IP:         "192.168.1.1",
		Mac:        "AA:BB:CC:DD:EE:01",
		Hw:         "Gateway Vendor",
		Date:       "2026-08-24 08:00:00",
		Known:      1,
		Now:        1,
		DeviceType: "router",
	})
	hosts := gdb.SelectByMAC("now", "AA:BB:CC:DD:EE:01")
	if len(hosts) != 1 {
		t.Fatalf("seeded hosts len = %d, want 1", len(hosts))
	}

	compareHosts(map[string]models.Host{
		"AA:BB:CC:DD:EE:01": {
			Iface: "wifi0",
			IP:    "192.168.1.254",
			Mac:   "AA:BB:CC:DD:EE:01",
			Hw:    "Scanned Gateway Vendor",
			Date:  "2026-08-24 09:00:00",
			Now:   1,
		},
	})

	updated := gdb.SelectByID(hosts[0].ID)
	if updated.DeviceType != "router" {
		t.Fatalf("DeviceType = %q, want router", updated.DeviceType)
	}
	if updated.Iface != "wifi0" || updated.IP != "192.168.1.254" || updated.Now != 1 {
		t.Fatalf("network-derived fields were not refreshed: %+v", updated)
	}
}

func TestNewHostCreatesOneDiscoveredEvent(t *testing.T) {
	setupScanRoutineTest(t)

	processScanResult([]models.Host{
		{
			Iface: "eth0",
			IP:    "127.0.0.1",
			Mac:   "AA:BB:CC:DD:EE:10",
			Hw:    "New Device Vendor",
			Date:  "2026-08-24 10:00:00",
			Now:   1,
		},
	}, true)

	events, ok := gdb.SelectEvents(10, "")
	if !ok {
		t.Fatal("SelectEvents failed")
	}
	if len(events) != 1 {
		t.Fatalf("events len = %d, want 1", len(events))
	}
	if events[0].EventType != string(models.EventDiscovered) {
		t.Fatalf("EventType = %q, want %q", events[0].EventType, models.EventDiscovered)
	}
	if events[0].Mac != "AA:BB:CC:DD:EE:10" {
		t.Fatalf("event Mac = %q, want AA:BB:CC:DD:EE:10", events[0].Mac)
	}
}

func TestRepeatedSuccessfulScanDoesNotDuplicateDiscoveredEvent(t *testing.T) {
	setupScanRoutineTest(t)

	host := models.Host{
		Iface: "eth0",
		IP:    "192.168.1.10",
		Mac:   "AA:BB:CC:DD:EE:10",
		Hw:    "New Device Vendor",
		Date:  "2026-08-24 10:00:00",
		Now:   1,
	}

	processScanResult([]models.Host{host}, true)
	host.Date = "2026-08-24 10:05:00"
	processScanResult([]models.Host{host}, true)

	events, ok := gdb.SelectEvents(10, "")
	if !ok {
		t.Fatal("SelectEvents failed")
	}
	if len(events) != 1 {
		t.Fatalf("events len = %d, want 1 discovered event: %+v", len(events), events)
	}
	if events[0].EventType != string(models.EventDiscovered) {
		t.Fatalf("EventType = %q, want %q", events[0].EventType, models.EventDiscovered)
	}
}

func TestOnlineOfflineTransitionsCreateSingleEvents(t *testing.T) {
	setupScanRoutineTest(t)

	gdb.Update("now", models.Host{
		Name:  "desktop",
		Iface: "eth0",
		IP:    "192.168.1.42",
		Mac:   "AA:BB:CC:DD:EE:42",
		Hw:    "Workstation NIC",
		Date:  "2026-08-24 08:00:00",
		Known: 1,
		Now:   0,
	})

	processScanResult([]models.Host{
		{
			Iface: "eth0",
			IP:    "192.168.1.42",
			Mac:   "AA:BB:CC:DD:EE:42",
			Hw:    "Workstation NIC",
			Date:  "2026-08-24 09:00:00",
			Now:   1,
		},
	}, true)
	assertEventTypes(t, []models.HostEventType{models.EventOnline})

	processScanResult([]models.Host{
		{
			Iface: "eth0",
			IP:    "192.168.1.42",
			Mac:   "AA:BB:CC:DD:EE:42",
			Hw:    "Workstation NIC",
			Date:  "2026-08-24 09:05:00",
			Now:   1,
		},
	}, true)
	assertEventTypes(t, []models.HostEventType{models.EventOnline})

	processScanResult([]models.Host{}, true)
	assertEventTypes(t, []models.HostEventType{models.EventOffline, models.EventOnline})

	processScanResult([]models.Host{}, true)
	assertEventTypes(t, []models.HostEventType{models.EventOffline, models.EventOnline})
}

func TestFailedScanDoesNotCreateOfflineEvent(t *testing.T) {
	setupScanRoutineTest(t)

	gdb.Update("now", models.Host{
		Name:  "router",
		Iface: "eth0",
		IP:    "192.168.1.1",
		Mac:   "AA:BB:CC:DD:EE:01",
		Hw:    "Gateway Vendor",
		Date:  "2026-08-24 08:00:00",
		Known: 1,
		Now:   1,
	})
	hosts := gdb.SelectByMAC("now", "AA:BB:CC:DD:EE:01")
	if len(hosts) != 1 {
		t.Fatalf("seeded hosts len = %d, want 1", len(hosts))
	}

	if processScanResult(nil, false) {
		t.Fatal("processScanResult returned true for failed scan")
	}

	updated := gdb.SelectByID(hosts[0].ID)
	if updated.Now != 1 {
		t.Fatalf("Now after failed scan = %d, want 1", updated.Now)
	}

	events, ok := gdb.SelectEvents(10, "")
	if !ok {
		t.Fatal("SelectEvents failed")
	}
	if len(events) != 0 {
		t.Fatalf("events len after failed scan = %d, want 0: %+v", len(events), events)
	}
}

func TestHostReturnsAfterFailedScanCreatesOnlyOnlineTransition(t *testing.T) {
	setupScanRoutineTest(t)

	gdb.Update("now", models.Host{
		Name:  "phone",
		Iface: "wifi0",
		IP:    "192.168.1.83",
		Mac:   "AA:BB:CC:DD:EE:83",
		Hw:    "Mobile Vendor",
		Date:  "2026-08-24 08:00:00",
		Known: 1,
		Now:   0,
	})

	if processScanResult(nil, false) {
		t.Fatal("processScanResult returned true for failed scan")
	}

	processScanResult([]models.Host{
		{
			Iface: "wifi0",
			IP:    "192.168.1.83",
			Mac:   "AA:BB:CC:DD:EE:83",
			Hw:    "Mobile Vendor",
			Date:  "2026-08-24 09:00:00",
			Now:   1,
		},
	}, true)

	assertEventTypes(t, []models.HostEventType{models.EventOnline})
}

func TestRetentionDeletesOldPresenceAndOnlyOldConnectivity(t *testing.T) {
	setupScanRoutineTest(t)

	host := models.Host{
		ID:    7,
		Name:  "storage",
		Iface: "eth0",
		IP:    "192.168.1.20",
		Mac:   "AA:BB:CC:DD:EE:20",
		Known: 1,
		Now:   1,
	}

	seededEvents := []struct {
		eventType models.HostEventType
		date      string
	}{
		{models.EventOnline, "2026-08-22 08:00:00"},
		{models.EventOffline, "2026-08-22 08:05:00"},
		{models.EventOnline, "2026-08-24 08:00:00"},
		{models.EventOffline, "2026-08-24 08:05:00"},
		{models.EventDiscovered, "2026-08-22 08:10:00"},
		{models.EventKnown, "2026-08-22 08:15:00"},
		{models.EventUnknown, "2026-08-22 08:20:00"},
		{models.EventDeviceTypeChanged, "2026-08-22 08:25:00"},
	}
	for _, item := range seededEvents {
		event := models.NewHostEvent(host, item.eventType, "", "")
		event.Date = item.date
		if err := gdb.AddEvent(event); err != nil {
			t.Fatalf("AddEvent %s: %v", item.eventType, err)
		}
	}

	gdb.Update("history", models.Host{
		Name: "old-presence-row",
		Mac:  "AA:BB:CC:DD:EE:20",
		Date: "2026-08-22 08:00:00",
	})
	gdb.Update("history", models.Host{
		Name: "recent-presence-row",
		Mac:  "AA:BB:CC:DD:EE:20",
		Date: "2026-08-24 08:00:00",
	})

	if deleted := gdb.DeleteOldHistory("2026-08-23 00:00:00"); deleted != 1 {
		t.Fatalf("DeleteOldHistory deleted = %d, want 1", deleted)
	}

	if deleted := gdb.DeleteOldConnectivityEvents("2026-08-23 00:00:00"); deleted != 2 {
		t.Fatalf("DeleteOldConnectivityEvents deleted = %d, want 2", deleted)
	}

	events, ok := gdb.SelectEvents(10, "")
	if !ok {
		t.Fatal("SelectEvents failed")
	}
	eventCounts := map[string]int{}
	for _, event := range events {
		eventCounts[event.EventType]++
		if (event.EventType == string(models.EventOnline) || event.EventType == string(models.EventOffline)) && event.Date < "2026-08-23 00:00:00" {
			t.Fatalf("old connectivity event remained: %+v", event)
		}
	}

	wantCounts := map[models.HostEventType]int{
		models.EventOnline:            1,
		models.EventOffline:           1,
		models.EventDiscovered:        1,
		models.EventKnown:             1,
		models.EventUnknown:           1,
		models.EventDeviceTypeChanged: 1,
	}
	for eventType, want := range wantCounts {
		if got := eventCounts[string(eventType)]; got != want {
			t.Fatalf("remaining %s events = %d, want %d; events: %+v", eventType, got, want, events)
		}
	}

	history, ok := gdb.Select("history")
	if !ok {
		t.Fatal("Select history failed")
	}
	if len(history) != 1 || history[0].Name != "recent-presence-row" {
		t.Fatalf("history rows = %+v, want only recent presence row", history)
	}
}

func assertEventTypes(t *testing.T, want []models.HostEventType) {
	t.Helper()

	events, ok := gdb.SelectEvents(10, "")
	if !ok {
		t.Fatal("SelectEvents failed")
	}
	if len(events) != len(want) {
		t.Fatalf("events len = %d, want %d: %+v", len(events), len(want), events)
	}
	for i, eventType := range want {
		if events[i].EventType != string(eventType) {
			t.Fatalf("events[%d].EventType = %q, want %q; events: %+v", i, events[i].EventType, eventType, events)
		}
	}
}
