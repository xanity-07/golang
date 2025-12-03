package intermediate

import (
	"fmt"
	"time"
)

func IntroTime() {
	//* Current local time
	now := time.Now()
	fmt.Println(now)

	//* Specific time
	specificTime := time.Date(2025, time.December, 3, 12, 20, 0, 0, time.UTC)
	fmt.Println(specificTime)

	//* Parse Time
	parsedTime, _ := time.Parse("2006-01-02", "2020-05-01") //? Mon Jan 2 15:04:05 MST 2006
	fmt.Println(parsedTime)

	//* Different format
	parseTime, _ := time.Parse("06-01-02", "20-05-01") //? Mon Jan 2 15:04:05 MST 2006
	fmt.Println(parseTime)
}
