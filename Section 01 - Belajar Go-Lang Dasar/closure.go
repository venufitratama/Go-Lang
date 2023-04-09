package main

import "fmt"

func main() {
	name := "Venu"
	counter := 0

	increment := func() {
		name := "Fitratama"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	fmt.Println("----")
	increment()
	increment()
	increment()

	fmt.Println("----")
	fmt.Println(counter)
	fmt.Println(name)
}
