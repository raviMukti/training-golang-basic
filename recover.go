package main

import "fmt"

// Recover will be catch Panic

func endApp() {
	errorMessage := recover()
	if errorMessage != nil {
		fmt.Println("Terjadi Error", errorMessage)
	}
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
