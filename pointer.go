package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// pass by value
	// address1 := Address{"Samarinda", "Kaltim", "Indo"} address2 := address1
	//
	// address2.City = "Grogot"
	// fmt.Println(address1)
	// fmt.Println(address2)

	// pass by reference
	address1 := Address{"Samarinda", "Kaltim", "Indo"}
	address2 := &address1

	address2.City = "Grogot"
	fmt.Println(address1) // ikut berubah
	fmt.Println(address2)

	fmt.Println()
	// pembahasan operator asterisk *
	// ini bakal memindahkan address2 ke
	// address2 = &Address{"Martapura", "Kalsel", " Indo"}
	// fmt.Println("address 1", address1)
	// fmt.Println("address 2", address2)

	fmt.Println("Dengan asterisk mengubah semua pointer")

	// mengubah value address2 dan siapapun yang mengacu kepadanya
	*address2 = Address{"Martapura", "Kalsel", " Indo"}
	fmt.Println("address 1", address1)
	fmt.Println("address 2", address2)
}
