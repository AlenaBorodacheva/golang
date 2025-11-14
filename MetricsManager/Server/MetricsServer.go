package main

import (
	metrics "MetricsServer/Metrics"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	// Регистрируем обработчик для пути "/"
	http.HandleFunc("/", handler)
	// Запускаем веб-сервер на порту 9999
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}

// Обработчик HTTP-запросов
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r.Method) // Тип метода
	fmt.Print(r.URL)    // запрашиваемый URL
	fmt.Print(r.Body)
	fmt.Println(r.Proto) // версия протокола

	if r.Method == http.MethodGet {
		curMetrics := metrics.GetMetrics()
		if curMetrics != nil {
			PrintInfo(curMetrics)
			data, err := json.Marshal(w)
			if err != nil {
				fmt.Println(err)
				return
			}
			w.Write([]byte(data))
		}
	}
}

func PrintInfo(metrics *metrics.Metrics) {
	fmt.Println("CPU:")
	fmt.Printf("ModelName: %+v\n", metrics.Cpu.ModelName)
	fmt.Printf("Cores: %+v\n", metrics.Cpu.Cores)
	fmt.Printf("Mhz: %+v\n", metrics.Cpu.Mhz)
	fmt.Printf("Percentages: %+v\n", metrics.Cpu.Percent)

	fmt.Println("RAM:")
	fmt.Printf("Total: %v bytes\n", metrics.Ram.Total)
	fmt.Printf("Free: %v bytes\n", metrics.Ram.Free)
	fmt.Printf("Used: %v bytes\n", metrics.Ram.Used)
	fmt.Printf("Used percent: %f%%\n", metrics.Ram.UsedPercent)
}
