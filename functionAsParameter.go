package main

import "fmt"

func sayHelloWithFilter(name string, filter func(string) string) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string {
	if name == "bangsat" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("ELghaz", spamFilter)
	sayHelloWithFilter("bangsat", spamFilter)
}
