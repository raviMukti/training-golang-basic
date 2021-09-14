package main

import "fmt"

func main() {

	// Constant tidak bisa diubah
	const jenisKelamin string = "PRIA"

	fmt.Println(jenisKelamin)

	// Tidak bisa diubah, ketika const tidak digunakan, apps tidak akan kena error compiler
	// jenisKelamin = "WANITA" ini akan error

	// Declare multiple constant at once
	const (
		pria   = "PRIA"
		wanita = "WANITA"
	)

	fmt.Println(pria)
	fmt.Println(wanita)

}
