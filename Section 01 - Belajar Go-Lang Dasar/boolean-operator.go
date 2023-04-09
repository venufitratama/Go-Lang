package main

import "fmt"

func main() {

	var ujian = 80
	var absensi = 80

	var lulusUjian = ujian >= 80
	var lulusAbsensi = absensi >= 80
	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsensi)

	var lulus = lulusUjian && lulusAbsensi
	fmt.Println(lulus)

	//namun, umunya akan langsung seperti ini:
	fmt.Println(ujian >= 80 && absensi >= 0)
}
