package main

// function panic digunakan untuk menghentikan total program
// tapi tidak dengan defer
// makanya kita bisa memanfaatkannya dengan fungsi recover()
// sehingga program tak akan berhenti karena panic karena
// recover nya di dalam defer
// recover menampung pesan panic dan balikin program btw
import "fmt"

func endApp() {
	fmt.Println("endApp")
	message := recover()
	fmt.Println("Terjadi panic", message)
}

func startApp(error bool) {
	defer endApp()
	if error {
		panic("Error cok")
	}
}

func main() {
	startApp(true)
	fmt.Println("Program masih jalan cuy")
}
