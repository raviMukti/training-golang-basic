package main

import "fmt"

func main() {

	// aritmatik
	var a = 10
	var b = 20
	var c = a + b

	fmt.Println(c)

	// Augmented Assignment
	var d = 10
	d += 10

	fmt.Println(d)

	// Unary Operator ++ -- - + !
	d++ // d + 1
	fmt.Println(d)

	var negative = -10
	var positive = +10

	fmt.Println(negative)
	fmt.Println(positive)
}
