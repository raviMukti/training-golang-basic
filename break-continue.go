package main

import "fmt"

func main() {

	for i := 0; i < 4; i++ {
		if i == 3 {
			break
		}
		fmt.Println("Perulangan Ke ", i)
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Angka Ganjil ", i)
	}
}
