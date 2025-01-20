package main

// kata kunci for bisa dipakai untuk for loop dan while
// loop dan juga bisa untuk for range/foreach btw foreach
// ini untuk iterasi terhadap data collection

import "fmt"

func main() {
	// contoh for untuk while loop

	fmt.Println("contoh for untuk while loop")
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}
	fmt.Println("Done while loop")

	// contoh untuk for loop bisa
	fmt.Println("contoh for untuk for loop biasa")
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}
	fmt.Println("Done for loop")

	fmt.Println("mengakses data collection loop manual")
	names := []string{"Ramohan", "EL", "Ghazali"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	fmt.Println("mengakses data collection dengan for range/foreach")

	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}

	fmt.Println("Jika tidak butuh index, ubah jadi garis bawah var pertama")
	fmt.Println("Karena di golang variabel yang dibuat wajib dipakai")
	fmt.Println("jika ga dipakai akan error")
	for _, name := range names {
		fmt.Println(name)
	}
}
