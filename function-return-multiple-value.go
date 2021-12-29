package main

import "fmt"

func getFullName() (string, string) {
	return "Ravi", "Mukti"
}

func main() {
	// Accessing Multiple Return Value
	// firstName, lastName := getFullName()
	// fmt.Println("First Name : "+firstName, "\nLast Name  : "+lastName)

	// Ignore Some Return Value With UnderScore
	firstName, _ := getFullName()
	fmt.Println("Only Just First Name : " + firstName)
}
