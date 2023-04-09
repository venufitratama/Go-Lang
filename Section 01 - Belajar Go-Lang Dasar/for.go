package main

import "fmt"

func main() {
	//counter := 1
	//
	////FOR LOOPS
	//for counter <= 10 {
	//	fmt.Println("Perulangan Ke-", counter)
	//	counter++ //simpel dari counter = counter + 1
	//}

	//FOR DENGAN STATEMENT

	//bentuk simpel dari kode diatas
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan Ke-", counter)
	}

	//FOR RANGE
	slice := []string{"Venu", "Fitratama", "S.M."}

	for i := 0; i < len(slice); i++ {
		//fmt.Println(i) //cek i dulu, kalo udh benar, lanjut
		fmt.Println(slice[i])
	}

	//FOR RANGE YANG LEBIH CEPAT
	for _, value := range slice {
		//fmt.Println("Index", i, "=", value) //ini dapat digunakan jika _ diganti dengan i
		fmt.Println(value)
	}

	person := make(map[string]string)
	person["name"] = "Venu Fitratama"
	person["title"] = "Programmer"
	person["age"] = "26"

	for key, value := range person { //map selalu menggunakan key, kalo slice/ array maka i
		fmt.Println(key, "=", value)
	}
}
