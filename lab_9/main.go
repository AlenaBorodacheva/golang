package main

import (
	"fmt"
	"os"
)

type Person struct {
	Name   string
	Age    int
	Status bool
	Values []int
}

func main() {
	person1 := Person{
		Name:   "Alex",
		Age:    25,
		Status: true,
		Values: []int{1, 2, 3},
	}
	person2 := Person{
		Name:   "Anna",
		Age:    35,
		Status: false,
		Values: []int{3, 4},
	}

	filename := "Persons.csv"
	os.Remove(filename)
	err1 := WriteToCsv(filename, person1)
	if err1 == nil {
		fmt.Println("Данные person1 успешно записаны в ", filename)
	}
	err2 := WriteToCsv(filename, person2)
	if err2 == nil {
		fmt.Println("Данные person2 успешно записаны в ", filename)
	}

	persons, err := ReadFromCsv(filename)
	if err == nil {
		fmt.Println("Данные успешно прочитаны из ", filename)
		for i := range persons {
			fmt.Printf("Person %d: %+v\n", i+1, persons[i])
		}
	}
}
