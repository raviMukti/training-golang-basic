package main

import "fmt"

func main() {

	// Convert dari int32 ke int64
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)

	var nilai16 int16 = int16(nilai32) // tidak sanggup convert karena nilainya sudah mencapai batas, dan akan kembali ke titik paling bawah dan terus sampe nilainya sama

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	var name = "Ravi"
	var r = name[0] // bentuknya byte

	var rString = string(r) // convert dari byte ke string

	fmt.Println(r)
	fmt.Println(rString)
}
