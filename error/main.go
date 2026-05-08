package main

import (
	"errors"
	"fmt"
)

var err error

func main() {

	result, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

}

func divide(a int, b int) (int, error) {

	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil

}
