package web

import (
	"log/slog"
	"time"

	"github.com/aceberg/WatchYourLAN/internal/db"
	"github.com/aceberg/WatchYourLAN/internal/models"
)

func trimHistoryRoutine() {

	for {
		trimHistory()

		time.Sleep(time.Duration(1) * time.Hour) // Every hour
	}
}

func trimHistory() {

	hours := appConfig.TrimHist

	nowStr := time.Now().Format("2006-01-02 15:04:05")  // This needed so all time is
	now, _ := time.Parse("2006-01-02 15:04:05", nowStr) // in one format
	nowMinus := now.Add(-time.Duration(hours) * time.Hour)

	if appConfig.HistInDB {
		histHosts = db.Select("history")
	}

	slog.Info("Removing all History before", "date", nowMinus.Format("2006-01-02 15:04:05"))

	n := 0
	newHistHosts := []models.Host{}
	ids := []int{}

	for _, hist := range histHosts {
		date, _ := time.Parse("2006-01-02 15:04:05", hist.Date)

		if date.Before(nowMinus) {
			n = n + 1
			if appConfig.HistInDB {
				ids = append(ids, hist.ID)
			}
		} else {
			newHistHosts = append(newHistHosts, hist)
		}
	}
	
	db.DeleteList(ids)
	histHosts = newHistHosts

	slog.Info("Removed records from History", "n", n)
}
