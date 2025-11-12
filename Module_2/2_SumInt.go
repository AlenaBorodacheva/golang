package main

func sumInt(numbers ...int) (count int, sum int) {
	sum = 0
	for _, number := range numbers {
		sum += number
	}
	count = len(numbers)
	return
}
