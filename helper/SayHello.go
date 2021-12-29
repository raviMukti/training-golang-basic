package helper

import "fmt"

// Function bisa diakses dari luar package karena diawali huruf kapital
func SayHello(name string) {
	fmt.Println("Hello " + name)
}

// Function tidak bisa diakses dari luar, karena diawali huruf kecil
func sayGoodBye(name string) {
	fmt.Println("GoodBye " + name)
}
