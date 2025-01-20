package main

import "fmt"

// defer untuk menempatkan eksekusi setelah function
// tempatnya dipanggil selesai dieksekusi

// fungsinya jika fungsi tempatnya dipanggil error
// sedangkan pemanggilan fungsi lainnya itu di bawah
// maka ga akan eksekusi
//
// jadi taroh defer di atas nya biar tetap dieksekusi, tetap
// dijalankan setelah function utama dieksekusi error atau ga
func logging() {
	fmt.Println("Selesai")
}

func runApp() {
	defer logging()
	fmt.Println("RunApp")
}

func main() {
	runApp()
}
