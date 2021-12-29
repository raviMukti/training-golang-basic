package main

import "fmt"

func main() {

	// Create Map
	var person = map[string]string{
		"nama":   "Ravi",
		"alamat": "Bandung",
	}

	fmt.Println(person)
	fmt.Println(person["nama"])
	fmt.Println(person["alamat"])

	// menambahkan atau merubah map
	person["title"] = "Programmer"
	fmt.Println(person)
	fmt.Println(person["title"])

	book := make(map[string]string)
	book["title"] = "Java"
	book["author"] = "Ravi"
	book["year"] = "2021"
	book["publisher"] = "KCN"

	fmt.Println(book)

	delete(book, "publisher")

	fmt.Println(book)

}
