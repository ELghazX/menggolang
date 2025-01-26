package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error ini")
	NotFoundError   = errors.New("not found error ini")
)

func GetByiD(id string) error {
	if id == "" {
		return ValidationError
	}
	if id != "ELghaz" {
		return NotFoundError
	}
	return nil
}

func main() {
	err := GetByiD("ELghaz")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("ValidationError")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("Not found")
		} else {
			fmt.Println("Unknown error")
		}
	}
}
