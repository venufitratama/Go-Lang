package main

import "fmt"

func logging() {
	fmt.Println("Selesai Memanggil Function")
}

//func runApplication() {
//	fmt.Println("Run Application")
//	logging()
//} //umumnya seperti ini

func runApplication(value int) {
	defer logging()
	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println("Result", result)
}

func main() {
	runApplication(0) //hasilnya error karena parameternya 0, namun func tetap terpanggil walaupun error.
}
