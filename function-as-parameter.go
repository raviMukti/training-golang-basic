package main

import "fmt"

// Function Type Declaration
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello " + filteredName)
}

// Regular Function As Parameter
// func sayHelloWithFilter(name string, filter func(string) string) {
// 	filteredName := filter(name)
// 	fmt.Println("Hello " + filteredName)
// }

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	// Call As A Return Value
	sayHelloWithFilter("Ravi", spamFilter) // Return Ravi

	// Call As A Return Value
	sayHelloWithFilter("Anjing", spamFilter) // Return ...

	// Call As A Value
	resultFilter := spamFilter
	sayHelloWithFilter("GoodWeather", resultFilter)
}
