package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("you blacklisted")
	} else {
		fmt.Println("Welcome")
	}
}

func main() {
	// function tanpa nama langsung di posisi value dimasukkan
	// ke variabel
	// kalau dah ada namanya tetap harus gaboleh parenthesis, tapi kalau langsung
	// gini boleh karena baru dibuat

	blacklist := func(name string) bool {
		return name == "bangsat"
	}

	registerUser("bangsat", blacklist)
}
