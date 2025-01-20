// map itu kek dictionary di python
// atau array asosiatif
// jumlah data dinamis
// key harus unique
package main

import "fmt"

func main() {
	mahasiswa := map[string]string{
		"name": "Muhammad Ghazali",
		"nim":  "1231247641",
	}

	fmt.Println("mahasiswa[name]", mahasiswa["name"])
	fmt.Println("mahasiswa[nim]", mahasiswa["nim"])

	fmt.Println(mahasiswa)
	delete(mahasiswa, "name")
	fmt.Println("setelah dihapus")
	fmt.Println(mahasiswa)

	novel := make(map[string]string)
	novel["title"] = "Senyum pertama di pagi airin"
}
