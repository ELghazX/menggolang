package main

import (
	"fmt"
)

func main() {
	// membuat variabel langsung inisialisasi nilai awal
	angka := 4
	// bisa disingkat begini
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

	// oke inii lebih jelasnya
	// untuk membuat variabel global di luar function manapun
	// membuat nya wajib seperti ini
	// var namaVar dataType = string
	var namaVar string
	var namaVar1 string = "Apapun"
	// tapi jika ada di dalam function termasuk func main
	// kita bisa singkat seperti ini jika ada data yang inisialisasi
	// namaVar := "Apa saja"
	// ingat ini harus di dalam suatu function atau variabel lokal
}
