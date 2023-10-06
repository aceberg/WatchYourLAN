package web

import (
	// "log"
	"strconv"
	"time"

	"github.com/aceberg/WatchYourLAN/internal/db"
)

func trimHistoryRoutine() {

	for {
		trimHistory()

		time.Sleep(time.Duration(1) * time.Hour)
	}
}

func trimHistory() {

	days, _ := strconv.Atoi(AppConfig.HistDays)
	now := time.Now()
	history := db.SelectHist(AppConfig.DbPath)

	for _, hist := range history {
		date, _ := time.Parse("2006-01-02 15:04:05", hist.Date)
		datePlus := date.Add(time.Duration(days) * 24 * time.Hour)

		if now.After(datePlus) {

			db.DeleteHist(AppConfig.DbPath, hist.ID)

			// log.Println("REMOVED DATE =", hist.Date)
		}
	}
}
