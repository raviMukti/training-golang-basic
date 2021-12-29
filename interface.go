package main

import "fmt"

type HasName interface {
	GetName() string
}

func doGreeting(name HasName) {
	fmt.Println("Hello", name.GetName())
}

type Person struct {
	Name string
}

// create a struct method that have a name like interface HasName
func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var person Person

	person.Name = "Ravi"

	doGreeting(person)

	cat := Animal{
		Name: "Tachi",
	}

	doGreeting(cat)
}
