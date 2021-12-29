package main

import "fmt"

func main() {
	// Closure is an ability to interract with data in it scope

	counter := 0

	increment := func() {
		fmt.Println("Increment")
		counter++ // It will override the counter above to be increment with 1
	}

	increment()
	increment() // Now counter will be 2

	fmt.Println(counter)

	// At least global scope, declare in above can be accessed with another funct below it
	// So what to do if we want to avoid override condition, make another variable name, and
	// re-declare the variable if we create a function
}
