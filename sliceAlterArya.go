package main

import "fmt"

func main() {
	// komponen slice :
	// pointer, length, capacity

	names := [...]string{"Arya", "Fayyadh", "Razan", "Muhammad", "Ghazali", "Ammar", "Nabil", "Fauzan"}
	slice := names[2:5]

	fmt.Println("slice[0]", slice[0])
	fmt.Println("slice[2]", slice[2])
	fmt.Println("len", len(slice))
	fmt.Println("cap", cap(slice))
	fmt.Println("append", append(slice, "Budiono"))

	//------------------------------------------
	//append yang melebih kapasitas array, akan membuat array baru
	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	fmt.Println(days)
	daySlice1 := days[5:]
	daySlice1[0] = "new sabtu"
	daySlice1[1] = "new minggu"

	fmt.Println(days)

	daySlice2 := append(daySlice1, "kiamat")
	daySlice2[0] = "ups"
	fmt.Println(daySlice2)
	fmt.Println(days)

	// ------------------------------------
	// make slice dipakai untuk membuat slice dari awal
	// tanpa array yang sudah ada
	// newSlice := make([]string, 2,5)
	// param1 = data type, param2 = length, param3 = capacity

	newSlice := make([]string, 2, 5)
	newSlice[0] = "ELghaz"
	newSlice[1] = "ELghaz"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Erugazu")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Budi"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	//copy slice
	//
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// hati hati saat membuat array bisa bisa malah jadi slice
	iniArya := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArya)
	fmt.Println(iniSlice)

	// pilih array atau slice?
}
