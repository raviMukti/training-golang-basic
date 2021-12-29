package main

import "fmt"

// If Use For Loop
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

// If Use Recursive Func
// Becareful with stackoverflow error u must set the stop condition, otherwise it's will throw an error
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	resultLoop := factorialLoop(3)
	fmt.Println(resultLoop)
	fmt.Println(factorialRecursive(3))
}
