package main

import (
	"fmt"
)

func swap() {
	a := 3
	b := 4
	fmt.Println("Value1", a)
	fmt.Println("Value2", b)

	p1 := &a
	p2 := &b
	fmt.Println("P1", p1)
	fmt.Println("P2", p2)

	*p1, *p2 = *p2, *p1

	fmt.Println("Value1 after swap", a)
	fmt.Println("Value2 after swap", b)

	fmt.Println("P1 after swap", p1)
	fmt.Println("P2 after swap", p2)
}
