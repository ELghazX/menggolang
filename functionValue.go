package main

// di golang function bisa dijadikan sebagai value

import "fmt"

func getByeByeYesterday(name string) string {
	return "Bye Bye " + name
}

func main() {
	// function yang sebagai value dimasukkan ke variabel.
	// variabel yang dibuat dengan value akan sama.
	// agar menjadi value dan tidak memanggil func yang
	// bersangkutan, jangan menggunakan parenthesis
	byeByeYesterday := getByeByeYesterday
	fmt.Println(byeByeYesterday("Yesterday"))
	contoh := byeByeYesterday

	fmt.Println(contoh("Contoh"))

	// dengan begini nantinya kita bisa mengiriimkan
	// function ke dalam parameter func lainnya
}
