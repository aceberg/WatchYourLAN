package models

import "testing"

func TestHostEventAutoMigrateCreatesEventsTable(t *testing.T) {
	db := openTempSQLite(t)

	if err := db.Table("events").AutoMigrate(&HostEvent{}); err != nil {
		t.Fatalf("AutoMigrate(events): %v", err)
	}

	for _, column := range []string{"ID", "HOST_ID", "MAC", "NAME", "EVENT_TYPE", "DATE", "IP", "IFACE", "DEVICE_TYPE", "OLD_VALUE", "NEW_VALUE"} {
		if !db.Table("events").Migrator().HasColumn(&HostEvent{}, column) {
			t.Fatalf("events table missing %s column after AutoMigrate", column)
		}
	}
}

func TestHostEventTypeValidation(t *testing.T) {
	for _, eventType := range HostEventTypeValues {
		if !IsValidHostEventType(string(eventType)) {
			t.Fatalf("event type %q should be valid", eventType)
		}
	}

	if IsValidHostEventType("renamed") {
		t.Fatal("renamed should not be a valid Phase 10 event type")
	}
}
