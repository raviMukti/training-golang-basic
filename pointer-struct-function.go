package main

import "fmt"

type Man struct {
	Name string
}

// Pass By Value tidak akan merubah apapun
// func (man Man) Married() {
// 	man.Name = "Mr. " + man.Name
// }

// Akan Merubah variable man
func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	person := Man{"Ravi"}
	person.Married()

	fmt.Println(person)
}
