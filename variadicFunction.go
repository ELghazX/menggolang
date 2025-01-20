// variadic function
// parameter yang berada di posisi terakhir bisa dijadikan varargs
// varargs bisa menerima lebih dari satu input, anggap aja kek array
// tapi bedanya dengan parameter tipe array
// gausah dibuat array nya dulu saat bikin argumen
package main

import "fmt"

// buatnya pakai 3 titik sebelum tipe data
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// jika udah ada slice dari awal
func main() {
	total := sumAll(1, 2, 3, 4, 5, 6)
	fmt.Println(total)
	// jika udah ada slice dari awal
	numbers := []int{10, 20, 30}
	fmt.Println(sumAll(numbers...))
}
