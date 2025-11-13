package main

import (
	"fmt"
	"os"
	"strconv"
)

func WriteSumToFile(sum int) {
	data := []byte(strconv.Itoa(sum))
	file, err := os.Create("sum.txt")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.Write(data)
}
