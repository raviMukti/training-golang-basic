package main

import "fmt"

func main() {
	// name := "Ravi"

	name := "Abed"

	if name == "Ravi" {
		fmt.Println("Hello Ravi")
	} else if name == "Abed" {
		fmt.Println("Hello Abed")
	} else {
		fmt.Println("Siapa Ya?")
	}

	// If Short Statement
	if nilai := 5; nilai > 4 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Sudah Benar")
	}
}
