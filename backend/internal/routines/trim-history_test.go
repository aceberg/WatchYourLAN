package routines

import (
	"testing"
	"time"

	"github.com/aceberg/WatchYourLAN/internal/conf"
	"github.com/aceberg/WatchYourLAN/internal/gdb"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

func TestTrimHistoryUsesIndependentPresenceAndConnectivityRetention(t *testing.T) {
	setupScanRoutineTest(t)
	conf.AppConfig.TrimHist = 48
	conf.AppConfig.ConnectivityRetention = 24

	host := models.Host{
		ID:    3,
		Name:  "desktop",
		Iface: "eth0",
		IP:    "192.168.1.42",
		Mac:   "AA:BB:CC:DD:EE:42",
		Known: 1,
		Now:   1,
	}
	seedPresenceHistory(t, "old-presence", "2026-08-22 11:59:59")
	seedPresenceHistory(t, "cutoff-presence", "2026-08-22 12:00:00")
	seedPresenceHistory(t, "new-presence", "2026-08-23 12:00:00")
	seedEvent(t, host, models.EventOnline, "2026-08-23 11:59:59")
	seedEvent(t, host, models.EventOffline, "2026-08-23 12:00:00")
	seedEvent(t, host, models.EventDiscovered, "2026-08-21 12:00:00")
	seedEvent(t, host, models.EventKnown, "2026-08-21 13:00:00")
	seedEvent(t, host, models.EventUnknown, "2026-08-21 14:00:00")
	seedEvent(t, host, models.EventDeviceTypeChanged, "2026-08-21 15:00:00")

	now := time.Date(2026, 8, 24, 12, 0, 0, 0, time.UTC)
	trimHistory(now)

	history, ok := gdb.Select("history")
	if !ok {
		t.Fatal("Select history failed")
	}
	if len(history) != 2 {
		t.Fatalf("history rows len = %d, want 2: %+v", len(history), history)
	}
	for _, row := range history {
		if row.Name == "old-presence" {
			t.Fatalf("old presence row survived independent retention cleanup: %+v", history)
		}
	}

	events, ok := gdb.SelectEvents(20, "")
	if !ok {
		t.Fatal("SelectEvents failed")
	}

	counts := map[string]int{}
	for _, event := range events {
		counts[event.EventType]++
		if event.EventType == string(models.EventOnline) {
			t.Fatalf("old online event survived connectivity retention cleanup: %+v", events)
		}
	}

	wantCounts := map[models.HostEventType]int{
		models.EventOffline:           1,
		models.EventDiscovered:        1,
		models.EventKnown:             1,
		models.EventUnknown:           1,
		models.EventDeviceTypeChanged: 1,
	}
	for eventType, want := range wantCounts {
		if got := counts[string(eventType)]; got != want {
			t.Fatalf("remaining %s events = %d, want %d; events: %+v", eventType, got, want, events)
		}
	}
}

func TestTrimHistoryFallsBackToPresenceRetentionWhenConnectivityRetentionMissing(t *testing.T) {
	setupScanRoutineTest(t)
	conf.AppConfig.TrimHist = 48
	conf.AppConfig.ConnectivityRetention = 0

	host := models.Host{
		ID:    4,
		Name:  "phone",
		Iface: "wifi0",
		IP:    "192.168.1.83",
		Mac:   "AA:BB:CC:DD:EE:83",
		Known: 1,
		Now:   1,
	}
	seedEvent(t, host, models.EventOnline, "2026-08-22 11:59:59")
	seedEvent(t, host, models.EventOffline, "2026-08-22 12:00:00")

	now := time.Date(2026, 8, 24, 12, 0, 0, 0, time.UTC)
	trimHistory(now)

	events, ok := gdb.SelectEvents(10, "")
	if !ok {
		t.Fatal("SelectEvents failed")
	}
	if len(events) != 1 {
		t.Fatalf("events len = %d, want 1: %+v", len(events), events)
	}
	if events[0].EventType != string(models.EventOffline) || events[0].Date != "2026-08-22 12:00:00" {
		t.Fatalf("remaining event = %+v, want cutoff offline event", events[0])
	}
}

func seedPresenceHistory(t *testing.T, name, date string) {
	t.Helper()

	gdb.Update("history", models.Host{
		Name: name,
		Mac:  "AA:BB:CC:DD:EE:42",
		Date: date,
	})
}

func seedEvent(t *testing.T, host models.Host, eventType models.HostEventType, date string) {
	t.Helper()

	event := models.NewHostEvent(host, eventType, "", "")
	event.Date = date
	if err := gdb.AddEvent(event); err != nil {
		t.Fatalf("AddEvent %s: %v", eventType, err)
	}
}
