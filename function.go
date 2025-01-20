package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

// function mengembalikan data/nilai
// tambahkan return dan tipe data sebelum
// buka kurung
func getHello(name string) string {
	return "Hello " + name
}

// di golang bisa mengembalikan multiple value
// dengan menuliskan type data yang ingin dikembalikan
// wrap dengan parenthesis ()

// misal mengembalikan nilai string dan int
func getFullMahasiswa() (string, int) {
	return "Muhammad Ghazali", 2309106041
}

// named return values
// membuat variabel untuk direturn
func getCompleteData() (firstName, lastName, nim string) {
	firstName = "Muhammad"
	lastName = "Ghazali"
	nim = "2309106041"
	return firstName, lastName, nim
}

func main() {
	sayHelloTo("Muhammad", "Ghazali")
	output := getHello("Elghaz")
	fmt.Println(output)

	// harus di unpack seperti di python
	outputNama, outputNim := getFullMahasiswa()
	fmt.Println(outputNama)
	fmt.Println(outputNim)

	fmt.Println("Langsung print ")
	fmt.Println(getFullMahasiswa())

	fmt.Println("di golang harus ditampung semua")
	fmt.Println("jika ingin menghiraukan salah satu return value")
	fmt.Println("sama seperti di foreach kita bissa pakai garis bawah")
	fmt.Println("misalkan hanya ingin mengambil nim saja")

	_, NIM := getFullMahasiswa()

	fmt.Println(NIM)

	a, b, c := getCompleteData()
	fmt.Println(a, b, c)
}
