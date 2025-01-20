package main

// function panic digunakan untuk menghentikan total program
// tapi tidak dengan defer
import "fmt"

func endApp() {
	fmt.Println("endApp")
}

func startApp(error bool) {
	defer endApp()
	if error {
		panic("Error cok")
	}
}

func main() {
	startApp(true)
}
