package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func MarshalUnmarshal(w Weather) {
	data, err := json.Marshal(w)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", &data)

	err = os.WriteFile("weather.json", data, 0644)
	if err != nil {
		fmt.Println("Ошибка записи в файл:", err)
		return
	}

	var weatherRead Weather
	fileData, err := os.ReadFile("weather.json")
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}
	json.Unmarshal(fileData, &weatherRead)
	fmt.Printf("%v", weatherRead)
}
