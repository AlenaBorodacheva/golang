package main

import (
	worker "MetricsClient/Worker"
	writer "MetricsClient/Writer"
	"time"
)

func main() {
	startPort := 7000
	endPort := 9999
	numWorkers := 10
	period := 1 * time.Minute

	db := writer.InitDB("Metrics.db")

	for {
		worker.Worker(startPort, endPort, numWorkers, db)
		time.Sleep(period)
	}
}
