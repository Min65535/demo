package week

import (
	"fmt"
	"time"
)

func GetSunday() {
	dateStr := "2022-01-01"
	date, _ := time.ParseInLocation("2006-01-02", dateStr, time.Local)
	fmt.Println("date:", date.Format("2006-01-02"))
	now := time.Now()
	now = date
	offset := int(time.Sunday - now.Weekday())
	if offset > 0 {
		offset = -6
	}

	weekStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	fmt.Println("weekStart:", weekStart.Format("2006-01-02"))
}
