package main

import (
	"errors"
	"fmt"
)

func divide(value int, dividen int) (int, error) {
	if dividen == 0 {
		return 0, errors.New("Divided By Zero")
	} else {
		result := value / dividen
		return result, nil
	}
}

func main() {
	// var exampleError error = errors.New("Oops Error")

	// fmt.Println(exampleError.Error())

	resultMath, err := divide(100, 10)
	if err == nil {
		fmt.Println("Hasil", resultMath)
	} else {
		fmt.Println("Error", err.Error())
	}
}
