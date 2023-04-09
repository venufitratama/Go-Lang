package main

import (
	"fmt"
)

func getCompleteName() (firstName, lastName, age string) { //bentuk singkat dari firstName string, lastName string, age string
	firstName = "Venu"
	lastName = "Fitratama"
	age = "26"
	return //firstName, lastName, age// ini tidak wajib karena sudah dideklarasikan diawal tadi
}

func main() {
	firstName, lastName, age := getCompleteName() //value bisa diganti dengan shift F6 2x > Action > Rename > Ganti nama
	fmt.Println(firstName, lastName, age)

	firstName, _, age = getCompleteName()
	fmt.Println(firstName, age)
}
