package main

import (
	"errors"
)

func divide(a, b int) (res int, err error) {
	if b == 0 {
		err = errors.New("ошибка")
	}

	res = a / b
	return
}
