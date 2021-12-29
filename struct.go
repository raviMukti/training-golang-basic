package main

import "fmt"

// Struct adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data
// lainnya dalam satu kesatuan
// Sederhananya Struct adalah kumpulan dari field
// Kaya bikin class di oop, di go kita buatnya struct

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	// Create var from struct
	var person Customer

	person.Name = "Ravi Mukti"
	person.Address = "Bandung"
	person.Age = 25

	fmt.Println(person)
	fmt.Println(person.Name)
	fmt.Println(person.Address)
	fmt.Println(person.Age)

	// Struct literals
	person2 := Customer{
		Name:    "Abed",
		Address: "Jakarta",
		Age:     24,
	}

	fmt.Println(person2)

	person3 := Customer{"John", "Jakarta", 25}

	fmt.Println(person3)
}
