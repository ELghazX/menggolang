package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	elghaz := Man{"ELghaz"}

	fmt.Println(elghaz.Name)
	elghaz.Married()
	fmt.Println(elghaz.Name)
}
