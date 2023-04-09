package main

import "fmt"

func getFullName() (string, string) {
	return "Venu", "Fitratama"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)

	//IGNORE BEBERAPA BAGIAN DARI RETURN VALUE
	//TINGGAL TAMBAHKAN "_" DI VARIABLE YANG INGIN DIGANTI
	firstName, _ = getFullName()
	fmt.Println(firstName)
}
