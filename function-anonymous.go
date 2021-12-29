package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blackList Blacklist) {
	if blackList(name) {
		fmt.Println("You Are Blocked " + name)
	} else {
		fmt.Println("Welcome " + name)
	}
}

// // If we choose to create a function with a name
// func blackListAdmin(name string) bool {
// 	return name == "admin"
// }

// func blackListRoot(name string) bool {
// 	return name == "root"
// }

func main() {
	// Anon as variable
	// blackList := func(name string) bool {
	// 	return name == "admin"
	// }

	// registerUser("admin", blackList)

	// Direct Declare
	registerUser("root", func(s string) bool {
		return s == "root"
	})

	registerUser("admin", func(s string) bool {
		return s == "root"
	})
}
