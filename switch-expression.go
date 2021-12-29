package main

import "fmt"

func main() {

	// name := "Ravi"

	name := "Nani"

	switch name {
	case "Ravi":
		fmt.Println("Hello Ravi")
	case "Abed":
		fmt.Println("Hello Abed")
	default:
		fmt.Println("Siapa Ya")
	}

	// switch short statement
	switch length := len(name); length > 3 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama Sudah Benar")
	}
}
