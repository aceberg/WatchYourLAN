package gdb

import (
	"log"
	"log/slog"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/aceberg/WatchYourLAN/internal/check"
	"github.com/aceberg/WatchYourLAN/internal/conf"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

var db *gorm.DB

// Start working with DB
func Start() {
	var tab *gorm.DB
	var err error
	var pgFail bool

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Warn,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)
	gormConf := &gorm.Config{Logger: newLogger}

	// Choose DB and connect
	if conf.AppConfig.UseDB == "postgres" {
		db, err = gorm.Open(postgres.Open(conf.AppConfig.PGConnect), gormConf)

		if err != nil {
			pgFail = true

			slog.Error("PostgreSQL connection error:", "err", err)
			slog.Warn("Falling back to SQLite")
		}
	}

	if pgFail || conf.AppConfig.UseDB != "postgres" {

		db, err = gorm.Open(sqlite.Open(conf.AppConfig.DBPath), gormConf)
		check.IfError(err)
	}

	// Migrate the schema
	tab = db.Table("now")
	err = tab.AutoMigrate(&models.Host{})
	check.IfError(err)

	tab = db.Table("history")
	err = tab.AutoMigrate(&models.Host{})
	check.IfError(err)
}
