package main

import "fmt"

func main() {
	username := "admin"
	password := "admin123"

	if username == "admin" && password == "admin1234" {
		fmt.Println("Berhasil login")
	} else {
		fmt.Println("login gagal")
	}
}
