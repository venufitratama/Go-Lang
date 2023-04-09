package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
	MaritalStatus bool
}

func main() {
	var venu Customer
	venu.Name = "Venu Fitratama"
	venu.Address = "Kubu Raya"
	venu.Age = 26
	venu.MaritalStatus = false

	fmt.Println(venu)
	fmt.Println(venu.Name)
	fmt.Println(venu.Address)
	fmt.Println(venu.Age)
	fmt.Println(venu.MaritalStatus)

	//cara lain bikin struct
	fio := Customer{
		Name:          "Fiola",
		Address:       "Pontianak",
		Age:           25,
		MaritalStatus: false,
	}
	fmt.Println(fio)

	//atau

	fiola := Customer{"Fiola", "Pontianak", 25, false}
	fmt.Println(fiola) //ini akan susah ketika struct diubah, rentan error.

}
