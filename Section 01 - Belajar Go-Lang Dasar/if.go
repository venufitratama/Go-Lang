package main

import "fmt"

func main() {
	name := "Venu"
	if name == "Venu" {
		fmt.Println("Hello Venu")
	}

	name2 := "Fitratama"
	if name2 == "Venu" {
		fmt.Println("Hello Venu") // ini tidak ada ke-print karna salah
	}

	//IF ELSE FUNCTION
	name3 := "Venu"
	if name3 == "Fitratama" {
		fmt.Println("Hello Venu")
	} else {
		fmt.Println("Who Are You?")
	}

	//IF ELIF FUNCTION
	name4 := "Venu"
	if name4 == "Fitratama" {
		fmt.Println("Hello Venu")
	} else if name4 == "Venu" {
		fmt.Println("Hi Venu")
	} else if name4 == "Venu--" {
		fmt.Println("Hi Venu--")
	} else {
		fmt.Println("Who Are You?")
	}

	//IF DENGAN SHORT STATEMENT
	//var length = len(name)
	//if length > 5 {
	//akan short jika begini
	if length := len(name); length > 5 {
		fmt.Println("Nama Terlalu Panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}
}
