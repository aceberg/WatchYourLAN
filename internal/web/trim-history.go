package web

import (
	"log"
	"strconv"
	"time"

	"github.com/aceberg/WatchYourLAN/internal/db"
)

func trimHistoryRoutine() {

	for {
		trimHistory()

		time.Sleep(time.Duration(1) * time.Hour) // Every hour
	}
}

func trimHistory() {

	days, _ := strconv.Atoi(AppConfig.HistDays)

	nowStr := time.Now().Format("2006-01-02 15:04:05")  // This needed so all time is
	now, _ := time.Parse("2006-01-02 15:04:05", nowStr) // in one format
	nowMinus := now.Add(-time.Duration(days) * 24 * time.Hour)

	history := db.SelectHist(AppConfig.DbPath)

	if AppConfig.LogLevel != "short" {
		log.Println("INFO: removing all history before", nowMinus.Format("2006-01-02 15:04:05"))
	}

	for _, hist := range history {
		date, _ := time.Parse("2006-01-02 15:04:05", hist.Date)

		if date.Before(nowMinus) {

			db.DeleteHist(AppConfig.DbPath, hist.ID)
		}
	}
}
