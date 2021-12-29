package main

import "fmt"

// Defer will be execute even the app thrown an error

func logging() {
	fmt.Println("Selesai Memanggil Function")
}

func runApplication(value int) {
	defer logging() // always use defer in the start
	fmt.Println("Run App")
	result := 10 / value
	fmt.Println("Result ", result)
	// logging() // if we use without defer
}

func main() {
	// runApplication(10) // it will be okay, logging func still call

	// will be thrown error, and loggin func never been called
	// runApplication(0)

	runApplication(10)
}
