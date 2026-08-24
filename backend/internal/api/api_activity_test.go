package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/aceberg/WatchYourLAN/internal/gdb"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

func TestActivityEndpointReturnsEmptyTable(t *testing.T) {
	router := setupTestRouter(t)

	rec := getPath(router, "/api/activity")
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body: %s", rec.Code, http.StatusOK, rec.Body.String())
	}

	var events []models.HostEvent
	if err := json.Unmarshal(rec.Body.Bytes(), &events); err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}
	if len(events) != 0 {
		t.Fatalf("events len = %d, want 0", len(events))
	}
}

func TestActivityEndpointDefaultAndExplicitLimit(t *testing.T) {
	router := setupTestRouter(t)
	host := seedHost(t, models.Host{Name: "router", Mac: "AA:BB:CC:DD:EE:01"})

	for i := 0; i < 25; i++ {
		event := models.NewHostEvent(host, models.EventOnline, "", "")
		event.Date = fmt.Sprintf("2026-08-24 10:%02d:00", i)
		if err := gdb.AddEvent(event); err != nil {
			t.Fatalf("AddEvent %d: %v", i, err)
		}
	}

	rec := getPath(router, "/api/activity")
	if rec.Code != http.StatusOK {
		t.Fatalf("default status = %d, want %d; body: %s", rec.Code, http.StatusOK, rec.Body.String())
	}
	if events := decodeActivityEvents(t, rec); len(events) != 20 {
		t.Fatalf("default events len = %d, want 20", len(events))
	}

	rec = getPath(router, "/api/activity?limit=2")
	if rec.Code != http.StatusOK {
		t.Fatalf("explicit status = %d, want %d; body: %s", rec.Code, http.StatusOK, rec.Body.String())
	}
	if events := decodeActivityEvents(t, rec); len(events) != 2 {
		t.Fatalf("explicit events len = %d, want 2", len(events))
	}

	rec = getPath(router, "/api/activity?limit=20")
	if rec.Code != http.StatusOK {
		t.Fatalf("compatible status = %d, want %d; body: %s", rec.Code, http.StatusOK, rec.Body.String())
	}
	if events := decodeActivityEvents(t, rec); len(events) != 20 {
		t.Fatalf("compatible events len = %d, want 20", len(events))
	}
}

func TestActivityEndpointRejectsInvalidLimit(t *testing.T) {
	router := setupTestRouter(t)

	tests := []string{
		"/api/activity?limit=abc",
		"/api/activity?limit=0",
		"/api/activity?limit=-1",
		"/api/activity?limit=101",
		"/api/host/1/activity?limit=abc",
		"/api/host/1/activity?limit=0",
		"/api/host/1/activity?limit=-1",
		"/api/host/1/activity?limit=101",
	}

	for _, path := range tests {
		t.Run(path, func(t *testing.T) {
			rec := getPath(router, path)
			if rec.Code != http.StatusBadRequest {
				t.Fatalf("status = %d, want %d; body: %s", rec.Code, http.StatusBadRequest, rec.Body.String())
			}
		})
	}
}

func TestActivityEndpointCategoryOmittedAndAllReturnEveryEventType(t *testing.T) {
	router := setupTestRouter(t)
	host := seedHost(t, models.Host{Name: "router", Mac: "AA:BB:CC:DD:EE:01"})
	seedActivityEventTypes(t, host)

	for _, path := range []string{"/api/activity?limit=10", "/api/activity?category=all&limit=10"} {
		t.Run(path, func(t *testing.T) {
			rec := getPath(router, path)
			if rec.Code != http.StatusOK {
				t.Fatalf("status = %d, want %d; body: %s", rec.Code, http.StatusOK, rec.Body.String())
			}
			events := decodeActivityEvents(t, rec)
			if len(events) != len(models.HostEventTypeValues) {
				t.Fatalf("events len = %d, want %d: %+v", len(events), len(models.HostEventTypeValues), events)
			}
		})
	}
}

func TestActivityEndpointFiltersByCategory(t *testing.T) {
	router := setupTestRouter(t)
	host := seedHost(t, models.Host{Name: "NAS", Mac: "AA:BB:CC:DD:EE:20"})
	seedActivityEventTypes(t, host)

	tests := []struct {
		path string
		want map[models.HostEventType]bool
	}{
		{
			path: "/api/activity?category=connectivity&limit=10",
			want: map[models.HostEventType]bool{
				models.EventOnline:  true,
				models.EventOffline: true,
			},
		},
		{
			path: "/api/activity?category=changes&limit=10",
			want: map[models.HostEventType]bool{
				models.EventDiscovered:        true,
				models.EventKnown:             true,
				models.EventUnknown:           true,
				models.EventDeviceTypeChanged: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			rec := getPath(router, tt.path)
			if rec.Code != http.StatusOK {
				t.Fatalf("status = %d, want %d; body: %s", rec.Code, http.StatusOK, rec.Body.String())
			}
			events := decodeActivityEvents(t, rec)
			if len(events) != len(tt.want) {
				t.Fatalf("events len = %d, want %d: %+v", len(events), len(tt.want), events)
			}
			for _, event := range events {
				if !tt.want[models.HostEventType(event.EventType)] {
					t.Fatalf("unexpected event type %q in %s: %+v", event.EventType, tt.path, events)
				}
			}
		})
	}
}

func TestActivityEndpointFiltersByEventType(t *testing.T) {
	router := setupTestRouter(t)
	host := seedHost(t, models.Host{Name: "router", Mac: "AA:BB:CC:DD:EE:01"})
	seedActivityEventTypes(t, host)

	tests := []struct {
		path string
		want map[models.HostEventType]bool
	}{
		{
			path: "/api/activity?eventType=online&limit=10",
			want: map[models.HostEventType]bool{
				models.EventOnline: true,
			},
		},
		{
			path: "/api/activity?eventType=known&eventType=unknown&limit=10",
			want: map[models.HostEventType]bool{
				models.EventKnown:   true,
				models.EventUnknown: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			rec := getPath(router, tt.path)
			if rec.Code != http.StatusOK {
				t.Fatalf("status = %d, want %d; body: %s", rec.Code, http.StatusOK, rec.Body.String())
			}
			events := decodeActivityEvents(t, rec)
			if len(events) != len(tt.want) {
				t.Fatalf("events len = %d, want %d: %+v", len(events), len(tt.want), events)
			}
			for _, event := range events {
				if !tt.want[models.HostEventType(event.EventType)] {
					t.Fatalf("unexpected event type %q in %s: %+v", event.EventType, tt.path, events)
				}
			}
		})
	}
}

func TestActivityEndpointRejectsInvalidEventType(t *testing.T) {
	router := setupTestRouter(t)

	for _, path := range []string{
		"/api/activity?eventType=",
		"/api/activity?eventType=ONLINE",
		"/api/activity?eventType=other",
	} {
		t.Run(path, func(t *testing.T) {
			rec := getPath(router, path)
			if rec.Code != http.StatusBadRequest {
				t.Fatalf("status = %d, want %d; body: %s", rec.Code, http.StatusBadRequest, rec.Body.String())
			}
		})
	}
}

func TestActivityEndpointRejectsInvalidCategory(t *testing.T) {
	router := setupTestRouter(t)

	for _, path := range []string{
		"/api/activity?category=",
		"/api/activity?category=online,offline",
		"/api/activity?category=ONLINE",
		"/api/activity?category=other",
	} {
		t.Run(path, func(t *testing.T) {
			rec := getPath(router, path)
			if rec.Code != http.StatusBadRequest {
				t.Fatalf("status = %d, want %d; body: %s", rec.Code, http.StatusBadRequest, rec.Body.String())
			}
		})
	}
}

func TestActivityEndpointIntersectsCategoryAndEventType(t *testing.T) {
	router := setupTestRouter(t)
	host := seedHost(t, models.Host{Name: "NAS", Mac: "AA:BB:CC:DD:EE:20"})
	seedActivityEventTypes(t, host)

	rec := getPath(router, "/api/activity?category=connectivity&eventType=online&eventType=known&limit=10")
	if rec.Code != http.StatusOK {
		t.Fatalf("intersection status = %d, want %d; body: %s", rec.Code, http.StatusOK, rec.Body.String())
	}
	events := decodeActivityEvents(t, rec)
	if len(events) != 1 || events[0].EventType != string(models.EventOnline) {
		t.Fatalf("intersection events = %+v, want only online", events)
	}

	rec = getPath(router, "/api/activity?category=connectivity&eventType=known&limit=10")
	if rec.Code != http.StatusOK {
		t.Fatalf("empty intersection status = %d, want %d; body: %s", rec.Code, http.StatusOK, rec.Body.String())
	}
	if events = decodeActivityEvents(t, rec); len(events) != 0 {
		t.Fatalf("empty intersection events = %+v, want none", events)
	}
}

func TestActivityEndpointOffsetDefaultAndExplicit(t *testing.T) {
	router := setupTestRouter(t)
	host := seedHost(t, models.Host{Name: "phone", Mac: "AA:BB:CC:DD:EE:83"})

	for i := 0; i < 5; i++ {
		event := models.NewHostEvent(host, models.EventOnline, "", "")
		event.Date = fmt.Sprintf("2026-08-24 10:%02d:00", i)
		if err := gdb.AddEvent(event); err != nil {
			t.Fatalf("AddEvent %d: %v", i, err)
		}
	}

	rec := getPath(router, "/api/activity?limit=2")
	if rec.Code != http.StatusOK {
		t.Fatalf("default offset status = %d, want %d; body: %s", rec.Code, http.StatusOK, rec.Body.String())
	}
	defaultEvents := decodeActivityEvents(t, rec)

	rec = getPath(router, "/api/activity?limit=2&offset=0")
	if rec.Code != http.StatusOK {
		t.Fatalf("explicit offset status = %d, want %d; body: %s", rec.Code, http.StatusOK, rec.Body.String())
	}
	offsetZeroEvents := decodeActivityEvents(t, rec)

	if len(defaultEvents) != len(offsetZeroEvents) || defaultEvents[0].Date != offsetZeroEvents[0].Date || defaultEvents[1].Date != offsetZeroEvents[1].Date {
		t.Fatalf("offset default events = %+v, want same as offset 0 %+v", defaultEvents, offsetZeroEvents)
	}

	rec = getPath(router, "/api/activity?limit=2&offset=2")
	if rec.Code != http.StatusOK {
		t.Fatalf("offset status = %d, want %d; body: %s", rec.Code, http.StatusOK, rec.Body.String())
	}
	offsetEvents := decodeActivityEvents(t, rec)
	if len(offsetEvents) != 2 || offsetEvents[0].Date != "2026-08-24 10:02:00" || offsetEvents[1].Date != "2026-08-24 10:01:00" {
		t.Fatalf("offset events = %+v, want third and fourth newest", offsetEvents)
	}
}

func TestActivityEndpointPaginatesWithEventTypeFilter(t *testing.T) {
	router := setupTestRouter(t)
	host := seedHost(t, models.Host{Name: "desktop", Mac: "AA:BB:CC:DD:EE:42"})

	for i := 0; i < 6; i++ {
		event := models.NewHostEvent(host, models.EventOnline, "", "")
		event.Date = fmt.Sprintf("2026-08-24 10:%02d:00", i)
		if err := gdb.AddEvent(event); err != nil {
			t.Fatalf("AddEvent online %d: %v", i, err)
		}
		change := models.NewHostEvent(host, models.EventKnown, "", "")
		change.Date = fmt.Sprintf("2026-08-24 09:%02d:00", i)
		if err := gdb.AddEvent(change); err != nil {
			t.Fatalf("AddEvent known %d: %v", i, err)
		}
	}

	rec := getPath(router, "/api/activity?eventType=online&limit=2&offset=2")
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body: %s", rec.Code, http.StatusOK, rec.Body.String())
	}
	events := decodeActivityEvents(t, rec)
	if len(events) != 2 || events[0].Date != "2026-08-24 10:03:00" || events[1].Date != "2026-08-24 10:02:00" {
		t.Fatalf("filtered page events = %+v, want third and fourth newest online events", events)
	}
	for _, event := range events {
		if event.EventType != string(models.EventOnline) {
			t.Fatalf("non-online event in filtered page: %+v", events)
		}
	}
}

func TestActivityEndpointRejectsInvalidOffset(t *testing.T) {
	router := setupTestRouter(t)

	for _, path := range []string{
		"/api/activity?offset=abc",
		"/api/activity?offset=-1",
	} {
		t.Run(path, func(t *testing.T) {
			rec := getPath(router, path)
			if rec.Code != http.StatusBadRequest {
				t.Fatalf("status = %d, want %d; body: %s", rec.Code, http.StatusBadRequest, rec.Body.String())
			}
		})
	}
}

func TestActivityEndpointCombinesCategoryOffsetAndLimit(t *testing.T) {
	router := setupTestRouter(t)
	host := seedHost(t, models.Host{Name: "router", Mac: "AA:BB:CC:DD:EE:01"})

	events := []struct {
		eventType models.HostEventType
		date      string
	}{
		{models.EventKnown, "2026-08-24 10:00:00"},
		{models.EventOnline, "2026-08-24 10:01:00"},
		{models.EventDeviceTypeChanged, "2026-08-24 10:02:00"},
		{models.EventOffline, "2026-08-24 10:03:00"},
		{models.EventUnknown, "2026-08-24 10:04:00"},
		{models.EventDiscovered, "2026-08-24 10:05:00"},
	}
	for _, item := range events {
		event := models.NewHostEvent(host, item.eventType, "", "")
		event.Date = item.date
		if err := gdb.AddEvent(event); err != nil {
			t.Fatalf("AddEvent: %v", err)
		}
	}

	rec := getPath(router, "/api/activity?category=changes&limit=2&offset=1")
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body: %s", rec.Code, http.StatusOK, rec.Body.String())
	}
	got := decodeActivityEvents(t, rec)
	want := []models.HostEventType{models.EventUnknown, models.EventDeviceTypeChanged}
	if len(got) != len(want) {
		t.Fatalf("events len = %d, want %d: %+v", len(got), len(want), got)
	}
	for i, eventType := range want {
		if got[i].EventType != string(eventType) {
			t.Fatalf("events[%d].EventType = %q, want %q; events: %+v", i, got[i].EventType, eventType, got)
		}
	}
}

func TestActivityEndpointNewestFirstByDateAndID(t *testing.T) {
	router := setupTestRouter(t)
	host := seedHost(t, models.Host{Name: "NAS", Mac: "AA:BB:CC:DD:EE:20"})

	first := models.NewHostEvent(host, models.EventDiscovered, "", "")
	first.Date = "2026-08-24 10:00:00"
	second := models.NewHostEvent(host, models.EventKnown, "", "")
	second.Date = "2026-08-24 10:00:00"
	third := models.NewHostEvent(host, models.EventOffline, "", "")
	third.Date = "2026-08-24 10:05:00"

	for _, event := range []models.HostEvent{first, second, third} {
		if err := gdb.AddEvent(event); err != nil {
			t.Fatalf("AddEvent: %v", err)
		}
	}

	rec := getPath(router, "/api/activity")
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body: %s", rec.Code, http.StatusOK, rec.Body.String())
	}

	events := decodeActivityEvents(t, rec)
	if len(events) != 3 {
		t.Fatalf("events len = %d, want 3", len(events))
	}
	want := []models.HostEventType{models.EventOffline, models.EventKnown, models.EventDiscovered}
	for i, eventType := range want {
		if events[i].EventType != string(eventType) {
			t.Fatalf("events[%d].EventType = %q, want %q; events: %+v", i, events[i].EventType, eventType, events)
		}
	}
}

func TestActivityEndpointFiltersByMacAndHostID(t *testing.T) {
	router := setupTestRouter(t)
	routerHost := seedHost(t, models.Host{Name: "router", Mac: "AA:BB:CC:DD:EE:01"})
	nasHost := seedHost(t, models.Host{Name: "NAS", Mac: "AA:BB:CC:DD:EE:20"})

	routerEvent := models.NewHostEvent(routerHost, models.EventOnline, "", "")
	routerEvent.Date = "2026-08-24 10:00:00"
	nasEvent := models.NewHostEvent(nasHost, models.EventOffline, "", "")
	nasEvent.Date = "2026-08-24 10:05:00"

	for _, event := range []models.HostEvent{routerEvent, nasEvent} {
		if err := gdb.AddEvent(event); err != nil {
			t.Fatalf("AddEvent: %v", err)
		}
	}

	rec := getPath(router, "/api/activity?mac="+routerHost.Mac)
	if rec.Code != http.StatusOK {
		t.Fatalf("mac filter status = %d, want %d; body: %s", rec.Code, http.StatusOK, rec.Body.String())
	}
	events := decodeActivityEvents(t, rec)
	if len(events) != 1 || events[0].Mac != routerHost.Mac {
		t.Fatalf("mac filter events = %+v, want only router event", events)
	}

	rec = getPath(router, "/api/host/"+itoa(nasHost.ID)+"/activity")
	if rec.Code != http.StatusOK {
		t.Fatalf("host filter status = %d, want %d; body: %s", rec.Code, http.StatusOK, rec.Body.String())
	}
	events = decodeActivityEvents(t, rec)
	if len(events) != 1 || events[0].HostID != nasHost.ID {
		t.Fatalf("host filter events = %+v, want only NAS event", events)
	}
}

func TestActivityEndpointFiltersByMultipleMacs(t *testing.T) {
	router := setupTestRouter(t)
	routerHost := seedHost(t, models.Host{Name: "router", Mac: "AA:BB:CC:DD:EE:01"})
	nasHost := seedHost(t, models.Host{Name: "NAS", Mac: "AA:BB:CC:DD:EE:20"})
	phoneHost := seedHost(t, models.Host{Name: "phone", Mac: "AA:BB:CC:DD:EE:83"})

	seedActivityEvent(t, routerHost, models.EventOnline, "2026-08-24 10:00:00")
	seedActivityEvent(t, nasHost, models.EventOffline, "2026-08-24 10:05:00")
	seedActivityEvent(t, phoneHost, models.EventKnown, "2026-08-24 10:10:00")

	query := "/api/activity?limit=10&mac=" + url.QueryEscape(routerHost.Mac) + "&mac=" + url.QueryEscape(nasHost.Mac)
	rec := getPath(router, query)
	if rec.Code != http.StatusOK {
		t.Fatalf("multi-mac status = %d, want %d; body: %s", rec.Code, http.StatusOK, rec.Body.String())
	}
	events := decodeActivityEvents(t, rec)
	if len(events) != 2 {
		t.Fatalf("multi-mac events len = %d, want 2: %+v", len(events), events)
	}
	for _, event := range events {
		if event.Mac != routerHost.Mac && event.Mac != nasHost.Mac {
			t.Fatalf("unexpected mac in multi-mac results: %+v", events)
		}
	}
}

func TestActivityEndpointCombinesCategoryAndMultipleMacs(t *testing.T) {
	router := setupTestRouter(t)
	routerHost := seedHost(t, models.Host{Name: "router", Mac: "AA:BB:CC:DD:EE:01"})
	nasHost := seedHost(t, models.Host{Name: "NAS", Mac: "AA:BB:CC:DD:EE:20"})
	phoneHost := seedHost(t, models.Host{Name: "phone", Mac: "AA:BB:CC:DD:EE:83"})

	seedActivityEvent(t, routerHost, models.EventOnline, "2026-08-24 10:00:00")
	seedActivityEvent(t, routerHost, models.EventKnown, "2026-08-24 10:01:00")
	seedActivityEvent(t, nasHost, models.EventOffline, "2026-08-24 10:05:00")
	seedActivityEvent(t, phoneHost, models.EventOnline, "2026-08-24 10:10:00")

	query := "/api/activity?category=connectivity&limit=10&mac=" + url.QueryEscape(routerHost.Mac) + "&mac=" + url.QueryEscape(nasHost.Mac)
	rec := getPath(router, query)
	if rec.Code != http.StatusOK {
		t.Fatalf("category multi-mac status = %d, want %d; body: %s", rec.Code, http.StatusOK, rec.Body.String())
	}
	events := decodeActivityEvents(t, rec)
	if len(events) != 2 {
		t.Fatalf("category multi-mac events len = %d, want 2: %+v", len(events), events)
	}
	for _, event := range events {
		if event.Mac != routerHost.Mac && event.Mac != nasHost.Mac {
			t.Fatalf("unexpected mac in category multi-mac results: %+v", events)
		}
		if event.EventType != string(models.EventOnline) && event.EventType != string(models.EventOffline) {
			t.Fatalf("unexpected event type in category multi-mac results: %+v", events)
		}
	}
}

func TestActivityStatsEndpointReturnsTotals(t *testing.T) {
	router := setupTestRouter(t)
	routerHost := seedHost(t, models.Host{Name: "router", Mac: "AA:BB:CC:DD:EE:01"})
	nasHost := seedHost(t, models.Host{Name: "NAS", Mac: "AA:BB:CC:DD:EE:20"})

	seedActivityEventTypes(t, routerHost)
	seedActivityEvent(t, nasHost, models.EventOnline, "2026-08-24 11:00:00")
	seedActivityEvent(t, nasHost, models.EventOffline, "2026-08-24 11:01:00")

	rec := getPath(router, "/api/activity/stats")
	if rec.Code != http.StatusOK {
		t.Fatalf("stats status = %d, want %d; body: %s", rec.Code, http.StatusOK, rec.Body.String())
	}
	stats := decodeActivityStats(t, rec)
	if stats.Total != 8 || stats.Online != 2 || stats.Offline != 2 || stats.Discovered != 1 || stats.Known != 1 || stats.Unknown != 1 || stats.DeviceTypeChanged != 1 {
		t.Fatalf("stats = %+v, want totals for all seeded events", stats)
	}
}

func TestActivityStatsEndpointAppliesDeviceFilterOnly(t *testing.T) {
	router := setupTestRouter(t)
	routerHost := seedHost(t, models.Host{Name: "router", Mac: "AA:BB:CC:DD:EE:01"})
	nasHost := seedHost(t, models.Host{Name: "NAS", Mac: "AA:BB:CC:DD:EE:20"})

	seedActivityEventTypes(t, routerHost)
	seedActivityEvent(t, nasHost, models.EventOnline, "2026-08-24 11:00:00")
	seedActivityEvent(t, nasHost, models.EventOffline, "2026-08-24 11:01:00")

	rec := getPath(router, "/api/activity/stats?mac="+url.QueryEscape(nasHost.Mac))
	if rec.Code != http.StatusOK {
		t.Fatalf("filtered stats status = %d, want %d; body: %s", rec.Code, http.StatusOK, rec.Body.String())
	}
	stats := decodeActivityStats(t, rec)
	if stats.Total != 2 || stats.Online != 1 || stats.Offline != 1 || stats.Discovered != 0 || stats.Known != 0 || stats.Unknown != 0 || stats.DeviceTypeChanged != 0 {
		t.Fatalf("filtered stats = %+v, want only NAS counts", stats)
	}
}

func TestActivityDevicesEndpointIncludesCurrentAndDeletedEventDevices(t *testing.T) {
	router := setupTestRouter(t)
	currentHost := seedHost(t, models.Host{
		Name:       "router",
		Mac:        "AA:BB:CC:DD:EE:01",
		DeviceType: "router",
	})
	quietHost := seedHost(t, models.Host{
		Name:       "quiet NAS",
		Mac:        "AA:BB:CC:DD:EE:20",
		DeviceType: "nas",
	})
	deletedHost := models.Host{
		ID:         99,
		Name:       "deleted tablet",
		Mac:        "AA:BB:CC:DD:EE:70",
		DeviceType: "tablet",
	}

	seedActivityEvent(t, currentHost, models.EventOnline, "2026-08-24 10:00:00")
	seedActivityEvent(t, deletedHost, models.EventOffline, "2026-08-24 11:00:00")

	rec := getPath(router, "/api/activity/devices")
	if rec.Code != http.StatusOK {
		t.Fatalf("devices status = %d, want %d; body: %s", rec.Code, http.StatusOK, rec.Body.String())
	}
	devices := decodeActivityDevices(t, rec)

	byMac := make(map[string]models.ActivityDeviceOption, len(devices))
	for _, device := range devices {
		byMac[device.Mac] = device
	}

	if device, ok := byMac[currentHost.Mac]; !ok || !device.Exists || device.HostID != currentHost.ID || device.DeviceType != "router" {
		t.Fatalf("current host option = %+v, ok=%v; devices=%+v", device, ok, devices)
	}
	if device, ok := byMac[quietHost.Mac]; !ok || !device.Exists || device.HostID != quietHost.ID || device.DeviceType != "nas" {
		t.Fatalf("quiet current host option = %+v, ok=%v; devices=%+v", device, ok, devices)
	}
	if device, ok := byMac[deletedHost.Mac]; !ok || device.Exists || device.HostID != deletedHost.ID || device.DeviceType != "tablet" {
		t.Fatalf("deleted event host option = %+v, ok=%v; devices=%+v", device, ok, devices)
	}
}

func decodeActivityEvents(t *testing.T, rec *httptest.ResponseRecorder) []models.HostEvent {
	t.Helper()

	var events []models.HostEvent
	if err := json.Unmarshal(rec.Body.Bytes(), &events); err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}
	return events
}

func decodeActivityStats(t *testing.T, rec *httptest.ResponseRecorder) models.ActivityStats {
	t.Helper()

	var stats models.ActivityStats
	if err := json.Unmarshal(rec.Body.Bytes(), &stats); err != nil {
		t.Fatalf("json.Unmarshal stats: %v", err)
	}
	return stats
}

func decodeActivityDevices(t *testing.T, rec *httptest.ResponseRecorder) []models.ActivityDeviceOption {
	t.Helper()

	var devices []models.ActivityDeviceOption
	if err := json.Unmarshal(rec.Body.Bytes(), &devices); err != nil {
		t.Fatalf("json.Unmarshal devices: %v", err)
	}
	return devices
}

func seedActivityEventTypes(t *testing.T, host models.Host) {
	t.Helper()

	for i, eventType := range models.HostEventTypeValues {
		seedActivityEvent(t, host, eventType, fmt.Sprintf("2026-08-24 10:%02d:00", i))
	}
}

func seedActivityEvent(t *testing.T, host models.Host, eventType models.HostEventType, date string) {
	t.Helper()

	event := models.NewHostEvent(host, eventType, "", "")
	event.Date = date
	if err := gdb.AddEvent(event); err != nil {
		t.Fatalf("AddEvent %s: %v", eventType, err)
	}
}
