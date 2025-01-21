package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// pass by value
	// address1 := Address{"Samarinda", "Kaltim", "Indo"} address2 := address1
	//
	// address2.City = "Grogot"
	// fmt.Println(address1)
	// fmt.Println(address2)

	// pass by reference
	address1 := Address{"Samarinda", "Kaltim", "Indo"}
	address2 := &address1

	address2.City = "Grogot"
	fmt.Println(address1)
	fmt.Println(address2)
}
