package main

import (
	"fmt"
	"menggolang/database"
	_ "menggolang/internal"
)

func main() {
	fmt.Println(database.GetDatabase())
}
