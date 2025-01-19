package main

import "fmt"

func main() {
	type nim string

	var idMahasiswa nim = "2309106041"
	fmt.Println(idMahasiswa)
	fmt.Println(nim("22222222"))

	var contoh string = "2309106040"
	var contohnim nim = nim(contoh)
	fmt.Println(contoh)
	fmt.Println(contohnim)
}
