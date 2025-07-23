package routines

import (
	"log/slog"
	"time"

	"github.com/aceberg/WatchYourLAN/internal/conf"
	"github.com/aceberg/WatchYourLAN/internal/gdb"
)

// HistoryTrim - routine for History
func HistoryTrim() {
	go func() {
		for {
			hours := conf.AppConfig.TrimHist
			nowMinus := time.Now().Add(-time.Duration(hours) * time.Hour)
			date := nowMinus.Format("2006-01-02 15:04:05")

			slog.Info("Removing all History before", "date", date)

			n := gdb.DeleteOldHistory(date)
			slog.Info("Removed records from History", "n", n)

			time.Sleep(time.Duration(1) * time.Hour) // Every hour
		}
	}()
}
