package main

import (
	"errors"
	"fmt"
)

func division(value, divideBy float32) (float32, error) {
	if divideBy == 0 {
		return 0, errors.New("tidak bisa dibagi 0")
	} else {
		return value / divideBy, nil
	}
}

func main() {
	result, err := division(5, 0)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println("Result", result)
	}
}
