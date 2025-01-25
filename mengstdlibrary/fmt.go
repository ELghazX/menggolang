package main

import "fmt"

func main() {
	firstName := "Muhammad"
	lastName := "Ghazali"

	fmt.Println("Hello '", firstName, lastName, "'")

	fmt.Printf("Hello '%s %s'\n", firstName, lastName)
}
