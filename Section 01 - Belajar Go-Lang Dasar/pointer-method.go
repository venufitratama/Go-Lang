package main

import "fmt"

type Man struct {
	Name string
}

//Jika menggunakan ini, tidak akan ada title pada nama
//func (man Man) Married() {
//	man.Name = "Mr. " + man.Name
//}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	venu := Man{"Venu"}
	venu.Married()

	fmt.Println(venu.Name)
}
