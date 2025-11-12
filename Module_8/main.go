package main

import (
	"fmt"
	"time"
)

func main() {
	w := weather{
		date: time.Now(),
	}
	fmt.Print(w)
}

type weather struct {
	date         time.Time `json:"date"`
	temperatureC int       `json:"temperatureC"`
	description  string    `json:"description"`
}
