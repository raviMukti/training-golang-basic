package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	// var person map[string]string = nil
	// fmt.Println(person)

	// var person map[string]string = newMap("")
	var person map[string]string = newMap("Ravi")

	if person == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(person)
	}
}
