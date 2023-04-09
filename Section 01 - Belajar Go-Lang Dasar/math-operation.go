package main

import "fmt"

func main() {

	//bisa langsung begini
	var result = 10 + 10
	fmt.Println(result)

	//atau begini
	var a = 10
	var b = 4
	var c = a + b
	fmt.Println(c)

	//Augmented Assignment
	var i = 10
	i += 10 // i = i + 10
	fmt.Println(i)

	//Unary Operator
	i++ // i = i + 1
	fmt.Println(i)

	var negative = -100
	var positive = 100
	fmt.Println(negative)
	fmt.Println(positive)
}
