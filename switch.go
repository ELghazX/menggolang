package main

import "fmt"

func main() {
	username := "admin1234"
	// password := "admin123"

	switch username {
	case "ELghaz":
		fmt.Println("ELghaz")
	case "admin1234":
		fmt.Println("admin1234")
	case "admin":
		fmt.Println("admin")
	default:
		fmt.Println("username tidak ada")
	}

	// note : switch case disini bisa pakai operator perbandingan
	// tapi mending pakai if else aja kalau gitu
}
