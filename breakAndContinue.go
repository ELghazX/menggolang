package main

import "fmt"

// break digunakan untuk menghentikan kode di bawahnya
// dan seluruh Perulangan sampai akhir
//
// continue digunakan untuk menghentikan perulangan saat ini
// dan di bawahnya tapi kemudian
// dan lanjut ke iterasi selanjutnya

func main() {
	for i := 0; i < 10; i++ {
		if i == 6 {
			break
		}
		fmt.Println("Perulangan ke - ", i)
	}
	fmt.Println()
	fmt.Println()
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Perulangan continue ke - ", i)
	}
}
