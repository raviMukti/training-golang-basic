package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

// Struct Method / Function
func (customer Customer) greeting(name string) {
	fmt.Println("Hello", name, "My Name Is", customer.Name)
}

func main() {
	var person Customer

	person.Name = "Ravi"
	person.Address = "Bandung"
	person.Age = 25

	person.greeting("John")
}
