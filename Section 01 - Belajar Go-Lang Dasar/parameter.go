package main

import "fmt"

func sayHelloTo(firstName string, lastName string, age int8) {
	fmt.Println("Hello", firstName, lastName, age)
}

func main() {
	sayHelloTo("Venu", "Fitratama", 26)
	//atau bisa begini
	firstName := "Fiola"
	sayHelloTo(firstName, "Okta", 25)
}
