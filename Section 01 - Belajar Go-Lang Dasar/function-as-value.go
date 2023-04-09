package main

import "fmt"

func getGoodBye(name string) string {
	return ("Goodbye ") + name
}

func main() {
	goodbye := getGoodBye("Venu")
	fmt.Println(goodbye)
}
