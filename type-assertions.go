package main

import "fmt"

func random() interface{} {
	// return "Ups"
	return true
}

func main() {
	// var result interface{} = random()

	// var resultString string = result.(string) // type conversion, if error occured panic will throw

	// fmt.Println(resultString)

	result := random()

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Integer", value)
	case bool:
		fmt.Println("Boolean", value)
	default:
		fmt.Println("Unknown Type")
	}
}
