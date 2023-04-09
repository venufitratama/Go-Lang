package main

import "fmt"

func sayHello() {
	fmt.Println("Hello World")
}

func repeat() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
}

func main() {
	//sayHello()
	//sayHello()
	//sayHello()
	for i := 0; i < 10; i++ {
		sayHello()
		repeat()
	}
}
