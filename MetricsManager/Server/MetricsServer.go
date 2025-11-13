package main

import (
	metrics "MetricsServer/Metrics"
	"fmt"
	"net/http"
)

func main() {
	// Регистрируем обработчик для пути "/"
	//http.HandleFunc("/", handler)
	// Запускаем веб-сервер на порту 9999
	//err := http.ListenAndServe(":9999", nil)
	//if err != nil {
	//	fmt.Println("Ошибка запуска сервера:", err)
	//}

	metrics.GetMetrics()
}

// Обработчик HTTP-запросов
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r.Method) // Тип метода
	fmt.Print(r.URL)    // запрашиваемый URL
	fmt.Print(r.Body)
	fmt.Println(r.Proto) // версия протокола
	w.Write([]byte("Привет!"))
	myMetrics := metrics.GetMetrics()
	fmt.Println(myMetrics)
}
