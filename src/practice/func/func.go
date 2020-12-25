package main

import (
	"errors"
	"fmt"
)

type operator func(x, y int) int

func calculate(x , y int, op operator) (int, error) {

	if op == nil {

		return 0, errors.New("invaild operator")
	}

	return op(x, y), nil
}
func main() {

	op := func(x, y int) int {

		return x + y
	}

	fmt.Println(calculate(1, 1, op))
}