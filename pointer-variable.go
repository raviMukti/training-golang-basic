package main

import "fmt"

// Pass By Value merupakan istilah untuk menduplikasi data di memory
// Pointer merupakan kemampuan membuat reference ke lokasi data di memory
// yang sama, tanpa menduplikasi data yang sudah ada
// sederhananya dengan kemampuan pointer kita bisa membuat pass by reference
// Pass By Reference biasanya menggunakan & didepan variable yg akan kita gunakan

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	// address2 := address1 // pass by value
	address2 := &address1 // pass by reference

	address2.City = "Bandung"

	// address2 = &Address{"Malang", "Jawa Timur", "Indonesia"} // address 1 tetap, address 2 berubah
	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"} // address 1 berubah, address 2 berubah

	fmt.Println(address1)
	fmt.Println(address2)

	var address3 *Address = new(Address)

	fmt.Println(address3)

	address3.City = "Bandung"
	address3.Province = "Jawa Barat"
	address3.Country = "Indonesia"

	fmt.Println(address3)
}
