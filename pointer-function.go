package main

import "fmt"

type Address struct {
	City, Province, Country string
}

// Dengan parameter seperti ini address tidak akan berubah
// func ChangeCountryToIndonesia(address Address) {
// 	address.Country = "Indonesia"
// }

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	var address = Address{
		City:     "Bandung",
		Province: "Jawa Barat",
		Country:  "",
	}

	// ChangeCountryToIndonesia(address) // address tidak berubah negaranya
	ChangeCountryToIndonesia(&address) // address akan berubah negaranya

	fmt.Println(address)
}
