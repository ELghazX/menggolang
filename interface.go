package main

// interface adalah sebuah kontrak
// tidak harus diimplementasikan secara explisit
// seperti di java oop
// public class BlaBla implements Fulan
//
// asalkan struct sudah sesuai kontrak interface maka akan dianggap
// sebagai implementasi dari interface

import "fmt"

// contoh interface
type HasName interface {
	GetName() string // dengan kontrak memiliki method getName() return string
}

func sayHello(hasName HasName) { // function dengan parameter interface
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string // sebuah contoh struct
}

func (person Person) GetName() string { // yang memiliki method GetaName() return string
	return person.Name
} // maka ini akan dianggap implementasi dari interface HasName

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{"ELghaz"}
	var animal Animal
	animal.Name = "Woof woof"

	sayHello(person)
	sayHello(animal)
}
