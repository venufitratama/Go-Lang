package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKTPvenu NoKTP = "**************2"
	var marriedStatus Married = false
	fmt.Println(noKTPvenu)
	fmt.Println(marriedStatus)
}
