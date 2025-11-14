package main

import (
	metrics "MetricsServer/Metrics"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	// Регистрируем обработчики
	http.HandleFunc("/ram", handlerRAM)
	http.HandleFunc("/cpu", handlerCPU)
	// Запускаем веб-сервер на порту 9999
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}

// Обработчик HTTP-запросов
func handlerRAM(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		ramMetrics := metrics.GetRam()
		if ramMetrics != nil {
			PrintRAM(ramMetrics)
			arrMetrics := [1]*metrics.RAM{ramMetrics}
			data, err := json.Marshal(arrMetrics)
			if err != nil {
				fmt.Println(err)
				return
			}
			w.Write([]byte(data))
			fmt.Println("Данные отправлены")
		}
	}
}

func handlerCPU(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		cpuMetrics := metrics.GetCpu()
		if cpuMetrics != nil {
			PrintCPU(cpuMetrics)
			arrMetrics := [1]*metrics.CPU{cpuMetrics}
			data, err := json.Marshal(arrMetrics)
			if err != nil {
				fmt.Println(err)
				return
			}
			w.Write([]byte(data))
			fmt.Println("Данные отправлены")
		}
	}
}

func PrintRAM(metrics *metrics.RAM) {
	fmt.Println("RAM:")
	fmt.Println("Date:", metrics.Date)
	fmt.Printf("Total: %v bytes\n", metrics.Total)
	fmt.Printf("Free: %v bytes\n", metrics.Free)
	fmt.Printf("Used: %v bytes\n", metrics.Used)
	fmt.Printf("Used percent: %f%%\n", metrics.UsedPercent)
}

func PrintCPU(metrics *metrics.CPU) {
	fmt.Println("CPU:")
	fmt.Println("Date:", metrics.Date)
	fmt.Printf("ModelName: %+v\n", metrics.ModelName)
	fmt.Printf("Cores: %+v\n", metrics.Cores)
	fmt.Printf("Mhz: %+v\n", metrics.Mhz)
	fmt.Printf("Percentages: %+v\n", metrics.Percent)
}
