package main

import "fmt"

func main() {

	//membuat array dengan manual
	var names [2]string
	//tidak bisa lebih dari 2 index karena sudah di set 2

	names[0] = "Venu"
	names[1] = "Fitratama"

	fmt.Println(names[0])
	fmt.Println(names[1])

	//memebuat array langsung
	var values = [3]int{
		90,
		80,
		95,
	}
	fmt.Println(values) //keseluruhan, atau bisa juga satu-satu seperti:
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	//len digunakan untuk mendapatkan informasi panjang array
	fmt.Println(len(values))
	fmt.Println(len(names))
	//atau

	var lagi [10]string
	fmt.Println(len(lagi))

	//array[index] untuk mendapatkan data di posisi index
	//array[index] = value untuk mengubah data di posisi index
}
