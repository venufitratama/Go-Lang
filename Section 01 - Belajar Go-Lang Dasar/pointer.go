package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Malaysia"
}

func main() {
	//PASS BY VALUE
	address1 := Address{"Pontianak", "Kalimantan Barat", "Indonesia"}
	address2 := address1

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println("-----")
	address2.City = "Singkawang"

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println("-----")

	//OPERATOR &
	address3 := Address{"Sanggau", "Kalimantan Barat", "Indonesia"}
	address4 := &address3

	fmt.Println(address3)
	fmt.Println(address4)
	fmt.Println("-----")
	address4.City = "Ngabang"

	fmt.Println(address3)
	fmt.Println(address4)
	fmt.Println("-----")

	//OPERATOR *
	address5 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address6 := &address5

	fmt.Println(address5)
	fmt.Println(address6)
	fmt.Println("-----")

	address6.City = "Bandung"

	address6 = &Address{"Malang", "Jawa Timur", "Indonesia"}
	//*address6 = Address{"Malang", "Jawa Timur", "Indonesia"}

	fmt.Println(address5)
	fmt.Println(address6)
	fmt.Println("-----")

	address6 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	//*address6 = Address{"Malang", "Jawa Timur", "Indonesia"}

	fmt.Println(address5)
	fmt.Println(address6)
	fmt.Println("-----")

	//Pointer to Function

	var alamat = Address{
		City:     "Subang",
		Province: "Jawa Barat",
		Country:  "",
	}

	ChangeCountryToIndonesia(&alamat)
	fmt.Println(alamat)

	//atau bisa seperti ini
	//var alamatPointer *Address = &alamat
	//ChangeCountryToIndonesia(alamatPointer)
	//fmt.Println(alamat)
}
