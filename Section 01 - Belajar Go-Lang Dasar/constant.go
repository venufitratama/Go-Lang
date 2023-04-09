package main

import "fmt"

//func main() {
//	const firstnm string = "Venu"
//	const lastnm = "Fitratama"
//	const birthDate = 1997

//	fmt.Println(firstnm)
//	fmt.Println(lastnm)
//	fmt.Println(birthDate)
//}

// Multiple constant
func main() {
	const (
		namadepan    string = "Venu"
		namabelakang        = "Fitratama"
		tanggallahir        = 1997
	)

	fmt.Println(namadepan)
	fmt.Println(tanggallahir)
}
