package main

import (
	"fmt"
)

func main() {
	var angka = 4
	name := "ELghaz"

	const (
		nameconst string = "Erugazu"
		numconst         = 2
	)

	var (
		firstname = "Muhammad"
		lastname  = "Ghazali"
	)
	fmt.Println(angka)
	fmt.Println(name)
	fmt.Println(firstname, lastname)
}
