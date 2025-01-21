package main

import "fmt"

func random() any {
	return "12"
}

func main() {
	result := random()
	// resultString := result.(string)

	// var resultInt int = result.(int)
	// fmt.Println(resultInt) // gabisa convert

	// cara amannya pakai switch expression

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("int", value)
	default:
		fmt.Println("Ga tau")

	}
}
