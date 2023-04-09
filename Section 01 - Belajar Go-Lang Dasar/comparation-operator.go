package main

import "fmt"

func main() {

	var name1 = "Venu"
	var name2 = "Venu"

	var result bool = name1 == name2
	fmt.Println(result)

	var name3 = "Budi"
	var name4 = "Eko"

	var result2 bool = name3 > name4
	fmt.Println(result2) //akan false karena B dibawah E

	fmt.Println("------")

	var value1 = 100
	var value2 = 200

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}
