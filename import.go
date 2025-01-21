package main

import (
	"fmt"
	"menggolang/helper"
)

func main() {
	result := helper.SayHello("Erugazu")
	fmt.Println(result)

	fmt.Println(helper.Application)
	fmt.Println(helper.SayHello())
}
