package main

import "fmt"

func random() interface{} {
	return "Test"
}

func main() {
	var result interface{} = random()
	//var resultString string = result.(string) //pastikan type datanya benar, agar tidak terjadinya panic
	//fmt.Println(resultString)

	//TYPE ASSERTIONS SWITCH
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "Is String")
	case int:
		fmt.Println("Value", value, "Is Integer")
	default:
		fmt.Println("Unknown Type")
	}
}
