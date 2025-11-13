package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func ReadWrite() {
	re := regexp.MustCompile(`[\x00-\x1F\x7F]`)
	file, err := os.Open("num.txt")
	if err != nil {
		fmt.Println("Unable to open file:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	arr := []int{}
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				return
			}
		}

		lineNum := re.ReplaceAllString(line, "")
		num, err := strconv.Atoi(lineNum)
		if err == nil {
			arr = append(arr, num)
			fmt.Print(line)
		} else {
			fmt.Println(err)
		}
	}

	sum := 0
	for i := range arr {
		sum = sum + arr[i]
	}
	write(sum)
}

func write(sum int) {
	data := []byte(strconv.Itoa(sum))
	file, err := os.Create("sum.txt")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.Write(data)
}
