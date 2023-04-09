package main

import "fmt"

type Customer2 struct {
	Name, Address string
	Age           int
	MaritalStatus bool
}

func sayHi(customer Customer2, name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func main() {
	var venu Customer2
	venu.Name = "Venu"
	venu.Address = "Jl. Usaha Berama"
	venu.Age = 26
	venu.MaritalStatus = false

	sayHi(venu, "Venu Fitratama")
}

//akan menjadi struct jika hanya seperti ini

// func(customer Customer2) sayHi(name string) {
//	fmt.Println("Hello", name, "My Name is", customer.Name)
//}
//
//func main() {
//	var venu Customer2
//	venu.Name = "Venu"
//	venu.Address = "Jl. Usaha Berama"
//	venu.Age = 26
//	venu.MaritalStatus = false
//
//	venu.sayHi("Venu Fitratama")
//}
