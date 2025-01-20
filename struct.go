package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

// cara buat method struct
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is ", customer.Name)
}

func main() {
	var jorip Customer

	jorip.Name = "Jorip FachGyaatt"
	jorip.Address = "Jepang"
	jorip.Age = 40
	fmt.Println(jorip)
	fmt.Println(jorip.Name)
	fmt.Println(jorip.Address)
	fmt.Println(jorip.Age)

	// datanya bisa dimasukkan langsung
	// seperti constructor saat object nya dibuat
	//
	arya := Customer{"Aryacool", "Samarinda", 17}
	fmt.Println(arya)

	arya.sayHello("jorip")
	jorip.sayHello("Arya")
}
