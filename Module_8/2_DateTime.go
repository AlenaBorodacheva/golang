package main

import (
	"fmt"
	"strings"
	"time"
)

func GetDuration(dates, format string) time.Duration {
	result := strings.Split(dates, ",")
	date1, err1 := time.Parse(format, result[0])
	date2, err2 := time.Parse(format, result[1])
	if err1 == nil && err2 == nil {
		if date1.Unix() > date2.Unix() {
			return date1.Sub(date2)
		} else {
			return date2.Sub(date1)
		}
	}
	var comerr error
	if err1 != nil {
		comerr = err1
	} else {
		comerr = err2
	}
	fmt.Println("Ошибка формата:", comerr)
	return 0
}
