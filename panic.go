package main

import "fmt"

// Panic will be stop the entire app if thrown

func endApp() {
	fmt.Println("Aplikasi Selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi Berjalan")
}

func main() {
	// runApp(false)
	runApp(true)
}
