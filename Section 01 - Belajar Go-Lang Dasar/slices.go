package main

import "fmt"

func main() {
	var months = [...]string{ //digunakan ... karena belum tahu berapa jumlah array-nya, atau bisa langsung deklarasi berapa banyak
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7] //bulan ke-4 sampai bulan ke-6
	fmt.Println("slice1   :", slice1)
	fmt.Println("len      :", len(slice1)) //4, 5, 6 => artinya lennya/ panjangnya adalah 3
	fmt.Println("capacity :", cap(slice1)) //dihitung dari 4 hingga 12, maka didapat 8

	//months[5] = "Diubah"
	//fmt.Println(slice1)  //hati2 dalam mengubah isi slice

	var slice2 = months[2:4]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Venu")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"

	fmt.Println(slice2)
	fmt.Println(months)

	var newSlice = make([]string, 2, 5)
	newSlice[0] = "Venu"
	newSlice[1] = "Fitratama"
	fmt.Println(newSlice)

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
