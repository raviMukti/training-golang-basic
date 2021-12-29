package main

import "fmt"

func main() {

	// Declare array manual
	var names [3]string

	names[0] = "Ravi"
	names[1] = "Mukti"
	names[2] = "Hartadi"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// Declare Array Langsung Assign Value
	var nilai = [3]int{70, 80, 90}

	fmt.Println(nilai)
	fmt.Println(len(nilai))
}
