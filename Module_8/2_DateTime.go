package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func GetDuration() time.Duration {
	layout := "2006-01-02"
	var str string
	file, err := os.Open("date.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	fmt.Fscanln(file, &str)

	result := strings.Split(str, ",")
	parsedTime1, err1 := time.Parse(layout, result[0])
	parsedTime2, err2 := time.Parse(layout, result[1])

	if err1 == nil && err2 == nil {
		var diff time.Duration
		if parsedTime1.Unix() > parsedTime2.Unix() {
			diff = parsedTime1.Sub(parsedTime2)
		} else {
			diff = parsedTime2.Sub(parsedTime1)
		}

		return diff
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
