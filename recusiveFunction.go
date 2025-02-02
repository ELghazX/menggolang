package main

import "fmt"

func factorial(value int) int {
	if value == 1 {
		return 1
	}
	return value * factorial(value-1)
}

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func main() {
	fmt.Println(factorial(3))
}
