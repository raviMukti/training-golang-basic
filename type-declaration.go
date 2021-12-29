package main

import (
	"fmt"
)

func main() {

	// Alias Type sama seperti byte atau yg lain

	type NOKTP string // tipe data baru NOKTP, tapi ini bentuknya string

	var ktpRavi NOKTP = "32312312039203913231323"

	fmt.Println(ktpRavi)

	type MARRIED bool

	var marriedStatus MARRIED = false

	fmt.Println(marriedStatus)
}
