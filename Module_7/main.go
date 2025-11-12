package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	go myChanFunc(ch, 6)
	newNum := <-ch
	fmt.Println("chan: ", newNum)

	firstChan := make(chan int, 1)
	firstChan <- 10
	secondChan := make(chan int, 1)
	//secondChan <- 10
	stopChan := make(chan struct{}, 1)
	//str := new(Str)
	//stopChan <- *str

	go func() {
		res := calculator(firstChan, secondChan, stopChan)
		if res != nil {
			fmt.Println("CalcChan: ", <-res)
		} else {
			fmt.Println("empty CalcChain")
		}
	}()
	time.Sleep(2 * time.Second)

	cap := 2
	ch1 := make(chan int, cap)
	ch1 <- 10
	ch1 <- 20
	ch2 := make(chan int, cap)
	ch2 <- 5
	ch2 <- 8
	chOut := make(chan int, cap)
	merge2Channels(Calc, ch1, ch2, chOut, cap)
}

func myChanFunc(ch chan int, n int) {
	ch <- n + 1
}

func calculator(firstChan, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	resChan := make(chan int, 1)
	select {
	case val := <-firstChan:
		resChan <- val * val
	case val := <-secondChan:
		resChan <- val * 3
	case <-stopChan:
		close(resChan)
	}

	return resChan
}

func merge2Channels(fn func(int) int, in1, in2 <-chan int, out chan<- int, n int) {
	var wg sync.WaitGroup

	work := func(fn func(int) int, in1, in2 <-chan int, out chan<- int, i int) {
		defer wg.Done()
		x1 := <-in1
		x2 := <-in2
		res := fn(x1) + fn(x2)
		out <- res
		fmt.Printf("Merge %d: ", i)
		fmt.Println(res)
	}

	wg.Add(n)
	for i := 0; i < n; i++ {
		go work(fn, in1, in2, out, i)
	}
	wg.Wait()
}

func Calc(a int) int {
	return a * 3
}

type Str struct{}
