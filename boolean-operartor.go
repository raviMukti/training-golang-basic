package main

import "fmt"

func main() {
	var ujian = 75
	var absensi = 90

	var lulusUjian = ujian >= 80
	var lulusAbsensi = absensi >= 70

	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsensi)

	var lulus = lulusUjian && lulusAbsensi

	fmt.Println(lulus)

	fmt.Println(ujian >= 75 && absensi >= 75)
}
