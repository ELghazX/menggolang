package main

import "fmt"

// tipe data paling tinggi itu interface {} (kosong)
// atau any
func ups() any {
	return 1
}

func main() {
	kosong := ups()

	fmt.Println(kosong)
}
