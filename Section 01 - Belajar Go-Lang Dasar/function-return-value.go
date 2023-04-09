package main

import "fmt"

func getHello(name string) string {
	//result := "Hello"
	//return result
	//prinsip simpelnya seperti diatas, namun bisa juga dibuat seperti ini:
	if name == "" {
		return "Hello, Who Are You?"
	} else {
		return "Hello " + name
	}
}

func main() {
	result := getHello("Venu")
	fmt.Println(result)
	fmt.Println("atau langsung seperti ini")
	fmt.Println(getHello(""))
}
