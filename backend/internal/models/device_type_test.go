package models

import (
	"path/filepath"
	"testing"

	sqlite "github.com/aceberg/gorm-sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func TestHostAutoMigrateAddsDeviceTypeToCleanSQLiteTables(t *testing.T) {
	db := openTempSQLite(t)

	for _, table := range []string{"now", "history"} {
		if err := db.Table(table).AutoMigrate(&Host{}); err != nil {
			t.Fatalf("AutoMigrate(%s): %v", table, err)
		}
		if !db.Table(table).Migrator().HasColumn(&Host{}, "DEVICE_TYPE") {
			t.Fatalf("%s table missing DEVICE_TYPE column after AutoMigrate", table)
		}
	}
}

func TestLegacyHostRowWithoutDeviceTypeMigratesAsUnassigned(t *testing.T) {
	db := openTempSQLite(t)

	if err := db.Exec(`
		CREATE TABLE now (
			ID integer PRIMARY KEY AUTOINCREMENT,
			NAME text,
			DNS text,
			IFACE text,
			IP text,
			MAC text,
			HW text,
			DATE text,
			KNOWN integer,
			NOW integer
		)
	`).Error; err != nil {
		t.Fatalf("create legacy table: %v", err)
	}

	if err := db.Exec(`
		INSERT INTO now (NAME, DNS, IFACE, IP, MAC, HW, DATE, KNOWN, NOW)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, "legacy", "", "eth0", "192.168.1.10", "AA:BB:CC:DD:EE:10", "Legacy Vendor", "2026-08-24 10:00:00", 1, 1).Error; err != nil {
		t.Fatalf("insert legacy row: %v", err)
	}

	if err := db.Table("now").AutoMigrate(&Host{}); err != nil {
		t.Fatalf("AutoMigrate legacy table: %v", err)
	}
	if !db.Table("now").Migrator().HasColumn(&Host{}, "DEVICE_TYPE") {
		t.Fatal("legacy now table missing DEVICE_TYPE column after AutoMigrate")
	}

	var host Host
	if err := db.Table("now").First(&host, 1).Error; err != nil {
		t.Fatalf("read migrated legacy host: %v", err)
	}
	if host.DeviceType != "" {
		t.Fatalf("DeviceType = %q, want empty string for Unassigned", host.DeviceType)
	}
	if host.Name != "legacy" || host.Mac != "AA:BB:CC:DD:EE:10" || host.Known != 1 || host.Now != 1 {
		t.Fatalf("legacy fields changed after migration: %+v", host)
	}
}

func openTempSQLite(t *testing.T) *gorm.DB {
	t.Helper()

	db, err := gorm.Open(sqlite.Open(filepath.Join(t.TempDir(), "migration-test.db")), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase: true,
		},
	})
	if err != nil {
		t.Fatalf("gorm.Open: %v", err)
	}

	t.Cleanup(func() {
		sqlDB, err := db.DB()
		if err != nil {
			t.Errorf("db.DB: %v", err)
			return
		}
		if err := sqlDB.Close(); err != nil {
			t.Errorf("sqlDB.Close: %v", err)
		}
	})

	return db
}
