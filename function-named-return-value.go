package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Ravi"
	middleName = "Mukti"
	lastName = "Hartadi"
	return // Just return, dont need to call every variable again
}

func main() {
	firstName, middleName, lastName := getCompleteName()

	fmt.Println("First Name  : " + firstName)
	fmt.Println("Middle Name : " + middleName)
	fmt.Println("Last Name   : " + lastName)
}
