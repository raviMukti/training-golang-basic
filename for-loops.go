package main

import "fmt"

func main() {

	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan Ke ", counter)
		counter++
	}

	fmt.Println()

	for i := 1; i <= 5; i++ {
		fmt.Println("Perulangan Ke ", i)
	}

	slice := []string{"Ravi", "Mukti", "Hartadi"}

	println()

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	println()

	// for range
	for i, value := range slice {
		fmt.Println("Index ", i, "=", value)
		// fmt.Println(value)
	}

}
