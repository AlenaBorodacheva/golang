package main

import (
	"fmt"
	"time"
)

func main() {
	weatherNow := Weather{
		Date:         time.Now(),
		TemperatureC: 15,
		Description:  "Ветрено, возможен дождь",
	}
	MarshalUnmarshal(weatherNow)
	fmt.Println()

	diff := GetDuration()
	fmt.Printf("Difference in hours: %.2f\n", diff.Hours())
	fmt.Printf("Difference in minutes: %.2f\n", diff.Minutes())
	fmt.Printf("Difference in seconds: %.2f\n", diff.Seconds())

	ReadWrite()
}

type Weather struct {
	Date         time.Time `json:"date"`
	TemperatureC int       `json:"temperatureC"`
	Description  string    `json:"description"`
}
