package gdb

import (
	"errors"
	"log"
	"log/slog"
	"os"
	"regexp"
	"sync"
	"time"

	sqlite "github.com/aceberg/gorm-sqlite"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/aceberg/WatchYourLAN/internal/check"
	"github.com/aceberg/WatchYourLAN/internal/conf"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

var db *gorm.DB
var dbMu sync.RWMutex

var errDBNotStarted = errors.New("database is not started")
var postgresURLPasswordPattern = regexp.MustCompile(`(?i)(postgres(?:ql)?://[^:\s/@]+:)[^@\s]+@`)
var postgresKeywordPasswordPattern = regexp.MustCompile(`(?i)(password\s*=\s*)[^\s]+`)

// Close closes the active database connection.
func Close() error {
	dbMu.Lock()
	defer dbMu.Unlock()

	if db == nil {
		return nil
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	err = sqlDB.Close()
	if err == nil {
		db = nil
	}
	return err
}

// Start working with DB
func Start() {
	check.IfError(StartErr())
}

// StartErr starts or replaces the active DB using the current app config.
func StartErr() error {
	return startWithConfig(conf.GetAppConfig(), true)
}

// Reconnect validates and migrates a new DB before swapping it into service.
func Reconnect(config models.Conf) error {
	return startWithConfig(config, false)
}

func startWithConfig(config models.Conf, allowPostgresFallback bool) error {
	candidate, err := openMigrated(config, allowPostgresFallback)
	if err != nil {
		return err
	}

	dbMu.Lock()
	oldDB := db
	db = candidate
	dbMu.Unlock()

	if err := closeDB(oldDB); err != nil {
		slog.Warn("Failed to close previous DB connection", "err", err)
	}

	return nil
}

func newGormConfig() *gorm.Config {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             5 * time.Second,
			LogLevel:                  logger.Warn,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)
	return &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase: true,
			// So upper case Columns could work in both PostgreSQL and SQLite
		},
	}
}

func openMigrated(config models.Conf, allowPostgresFallback bool) (*gorm.DB, error) {
	candidate, err := connect(config, newGormConfig(), allowPostgresFallback)
	if err != nil {
		return nil, err
	}

	if err := migrate(candidate); err != nil {
		_ = closeDB(candidate)
		return nil, err
	}

	return candidate, nil
}

func migrate(candidate *gorm.DB) error {
	if err := candidate.Table("now").AutoMigrate(&models.Host{}); err != nil {
		return err
	}
	if err := candidate.Table("history").AutoMigrate(&models.Host{}); err != nil {
		return err
	}
	return candidate.Table("events").AutoMigrate(&models.HostEvent{})
}

func closeDB(target *gorm.DB) error {
	if target == nil {
		return nil
	}

	sqlDB, err := target.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}

func acquireDB() (*gorm.DB, func(), error) {
	dbMu.RLock()
	if db == nil {
		dbMu.RUnlock()
		return nil, func() {}, errDBNotStarted
	}

	return db, dbMu.RUnlock, nil
}

// Connect - choose DB and connect
func Connect() {
	candidate, err := connect(conf.GetAppConfig(), newGormConfig(), true)
	check.IfError(err)

	dbMu.Lock()
	oldDB := db
	db = candidate
	dbMu.Unlock()

	check.IfError(closeDB(oldDB))
}

func connect(config models.Conf, gormConf *gorm.Config, allowPostgresFallback bool) (*gorm.DB, error) {
	var pgFail bool
	var err error
	var candidate *gorm.DB

	if config.UseDB == "postgres" {
		candidate, err = gorm.Open(postgres.Open(config.PGConnect), gormConf)

		if err != nil {
			if !allowPostgresFallback {
				slog.Error("PostgreSQL connection error", "err", redactDatabaseError(err))
				return nil, err
			}
			pgFail = true

			slog.Error("PostgreSQL connection error", "err", redactDatabaseError(err))
			slog.Warn("Falling back to SQLite")
		} else {
			slog.Info("Connected to DB: PostgreSQL")
		}
	}

	if pgFail || config.UseDB != "postgres" {

		candidate, err = gorm.Open(sqlite.Open(config.DBPath), gormConf)

		if err != nil {
			return nil, err
		}
		slog.Info("Connected to DB: SQLite")
		candidate.Exec("PRAGMA journal_mode = wal;")
		candidate.Exec("PRAGMA busy_timeout = 5000;")
	}

	return candidate, nil
}

func redactDatabaseError(err error) string {
	return redactPostgresURL(err.Error())
}

func redactPostgresURL(value string) string {
	if value == "" {
		return ""
	}

	value = postgresURLPasswordPattern.ReplaceAllString(value, "${1}<redacted>@")
	return postgresKeywordPasswordPattern.ReplaceAllString(value, "${1}<redacted>")
}
