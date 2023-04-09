package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Venu",
		"address": "Jl. Pramuka",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	//boleh menambahkan key baru
	person["title"] = "Programmer"
	person["age"] = "26"
	fmt.Println(person)

	// len(map) untuk mendapatkan jumlah data Map
	// map[key] untuk mengambil data di map dengan key
	// map[key] = value untuk mengubah data di map dengan key
	// make(map[TypeKey]TypeValue) untuk membuat map baru
	// delete(map, key) untuk menghapus data di map dengan key

	fmt.Println("-----")
	fmt.Println(len(person))

	fmt.Println(person["age"])

	person["age"] = "26,5"
	fmt.Println(person["age"])

	book := make(map[string]string)
	book["title"] = "Go-Lang Introduction"
	book["Author"] = "Venu Fitratama"
	book["Wrong"] = "Delete Soon"
	fmt.Println(book)
	fmt.Println(len(book))
	delete(book, "Wrong")
	fmt.Println(book)
	fmt.Println(len(book))
}
