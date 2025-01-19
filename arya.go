package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Ramohan"
	names[1] = "EL"
	names[2] = "Ghazali"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(names)

	array := [5]int{
		1, 2, 3, 3, 3,
	}

	arya := [...]int{
		2, 3, 0, 9, 1, 0, 6, 0, 4, 1,
	}

	fmt.Println(array)
	fmt.Println(len(array))
	fmt.Println(arya)
	fmt.Println(len(arya))
}
