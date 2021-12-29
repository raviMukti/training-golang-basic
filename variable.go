package main

import "fmt"

func main() {
	// declare var dengan tipe data
	var name string

	name = "Ravi Mukti"

	fmt.Println(name)

	name = "Budi Sudarsono"

	fmt.Println(name)

	// declare variable tanpa tipe data
	var pekerjaan = "Backend"

	fmt.Println(pekerjaan)

	// declare variable tanpa menggunakan var
	alamat := "Bandung"

	fmt.Println(alamat)

	// Buat variable secara bersamaan atau multiple
	var (
		firstName = "Ravi"
		lastName  = "Mukti"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
