package main

import (
	"fmt"
)

func main() {
	fmt.Println("Fibonacci 10:", fibonacci(10))

	count, sum := sumInt(5, 8, 4)
	fmt.Println("Count: ", count)
	fmt.Println("Sum: ", sum)

	a := 3
	b := 4
	fmt.Println("Multiply: ", multiply(&a, &b))

	swap()

	res, err := divide(10, 2)
	fmt.Println(res, err)

	fmt.Println("valid password", valid("qwerty123"))
	fmt.Println("not valid password", valid("qwe"))
	fmt.Println("not valid password", valid("qцукен"))
}

func fibonacci(n uint) uint {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return fibonacci(n-1) + fibonacci(n-2)
}
