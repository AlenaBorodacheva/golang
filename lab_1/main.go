package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	defaultData := []int{64, 34, 25, 12, 22, 11, 90}
	array := []int{}

	fmt.Println("Введите элементы массива через запятую.")
	var arrayString string
	fmt.Scan(&arrayString)

	if arrayString == "" {
		array = defaultData
	} else {
		result := strings.Split(arrayString, ",")
		for i := range result {
			num, err := strconv.Atoi(result[i])
			if err != nil {
				array = defaultData
				break
			}
			array = append(array, num)
		}
	}

	fmt.Println("До сортировки:", array)
	MergeSort(array)
	fmt.Println("После сортировки:", array)
}
