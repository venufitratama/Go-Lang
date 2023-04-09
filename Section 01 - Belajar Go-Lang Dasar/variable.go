package main

import "fmt"

// variable cara 1
func main() {
	var name string

	name = "Venu Fitratama"
	fmt.Println(name)

	name = "Venu Fitratama--"
	fmt.Println(name) //akan berubah setelah update

	// variable dalam Go akan error jika tidak digunakan.

	//variable cara 2

	var birth = 97 //tanpa menuliskan type data > otomatis kedetect
	fmt.Println(birth)

	//variable 3
	namalengkap := "Venu" //tidak perlu menuliskan var
	fmt.Println(namalengkap)
	// := hanya digunakan untuk deklarasi awal, setelahnya tidak perlu digunakan

	//variable 4, kompilasi
	var (
		firstName = "venu"
		lastName  = "Fitratama"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}
