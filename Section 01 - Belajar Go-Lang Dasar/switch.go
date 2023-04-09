package main

import "fmt"

func main() {
	name := "Venu"

	//SWITCH DENGAN KONDISI
	switch name {
	case "Venu":
		fmt.Println("Hello Venu")
		fmt.Println("How Are You?")
	case "Fiola":
		fmt.Println("Hello Fio")
		fmt.Println("How Are You Today?")
		fmt.Println("Hope You Are Doing Great")
	default:
		fmt.Println("Who Are You?")
	}

	//SHORT STATEMENT
	//switch length := len(name); length > 4 {
	//case true:
	//	fmt.Println("Nama Terlalu Panjang")
	//case false:
	//	fmt.Println("Nama Sudah Benar")
	//}

	//SWITCH TANPA KONDISI
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama Terlalu Panjang")
	case length > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama Sudah Benar")
	}
}
