package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	startPort := 1
	endPort := 1000
	numWorkers := 10

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for range numWorkers {
		go worker(jobs, results, &wg)
	}

	go func() {
		defer close(jobs)
		for port := startPort; port <= endPort; port++ {
			jobs <- port
		}
	}()

	wg.Wait()
	close(results)

	fmt.Println("Открытые порты:")
	for port := range results {
		fmt.Println(port)
	}
}

func worker(jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for port := range jobs { // Читаем порты из канала jobs
		address := fmt.Sprintf("localhost:%d", port) // Проверяем на localhost
		conn, err := net.Dial("tcp", address)
		if err == nil { // Если соединение успешно — порт открыт
			conn.Close()    // Закрываем соединение
			results <- port // Отправляем номер порта в канал результатов
		}
	}
}
