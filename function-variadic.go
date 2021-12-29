package main

import "fmt"

// Variadic indicates with ... and its a slice
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	total := sumAll(10, 15, 15)
	fmt.Println("Total", total)

	// If U Want Slice to become a parameter for variadic function
	slice := []int{19, 123, 13, 12, 11}
	grandTotal := sumAll(slice...)
	fmt.Println("GrandTotal", grandTotal)
}
