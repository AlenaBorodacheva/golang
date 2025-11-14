package worker

import (
	writer "MetricsClient/Writer"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"

	"gorm.io/gorm"
)

func Worker(startPort, endPort, numWorkers int, db *gorm.DB) {
	jobs := make(chan int, 100)

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for range numWorkers {
		go worker(jobs, &wg, db)
	}

	go func() {
		defer close(jobs)
		for port := startPort; port <= endPort; port++ {
			jobs <- port
		}
	}()

	wg.Wait()
}

func worker(jobs <-chan int, wg *sync.WaitGroup, db *gorm.DB) {
	defer wg.Done()

	for port := range jobs { // Читаем порты из канала jobs
		addressRam := "http://localhost:" + strconv.Itoa(port) + "/ram"

		dataRam, err := getMetrics(addressRam)
		if err != nil {
			continue
		}

		if dataRam != nil {
			writer.WriteToCsv("RAM_"+strconv.Itoa(port)+".csv", dataRam)
			writer.WriteRamToSql(db, dataRam, port)
		}

		addressCpu := "http://localhost:" + strconv.Itoa(port) + "/cpu"
		dataCpu, _ := getMetrics(addressCpu)

		if dataCpu != nil {
			writer.WriteToCsv("CPU_"+strconv.Itoa(port)+".csv", dataCpu)
			writer.WriteCpuToSql(db, dataCpu, port)
		}
	}
}

func getMetrics(address string) ([]byte, error) {
	resp, err := http.Get(address)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()            // закрываем тело ответа после работы с ним
	data, err := io.ReadAll(resp.Body) // читаем ответ
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return data, nil
}
