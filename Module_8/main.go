package main

import (
	"fmt"
	"os"
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

	var dates string
	fmt.Print("Введите даты через запятую: ")
	fmt.Fscan(os.Stdin, &dates)
	diff := GetDuration(dates, "2020-02-01")
	fmt.Printf("Difference in hours: %.2f\n", diff.Hours())
	fmt.Printf("Difference in minutes: %.2f\n", diff.Minutes())
	fmt.Printf("Difference in seconds: %.2f\n", diff.Seconds())
	fmt.Println()

	fmt.Println("Введите числа от 0 до 100: ")
	sum := 0
	for {
		var num int
		_, err := fmt.Fscan(os.Stdin, &num)
		if err != nil {
			break
		}
		fmt.Print("Введено число ")
		fmt.Fprintln(os.Stdout, num)
		sum += num
	}
	fmt.Println("Сумма:", sum)
	WriteSumToFile(sum)
}

type Weather struct {
	Date         time.Time `json:"date"`
	TemperatureC int       `json:"temperatureC"`
	Description  string    `json:"description"`
}
