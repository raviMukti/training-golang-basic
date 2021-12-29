package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Ravi")
	data.PushBack("Mukti")
	data.PushBack("Hartadi")

	// fmt.Println(data.Front().Value) // Ambil data paling depan
	// fmt.Println(data.Back().Value)  // Ambil data paling belakang

	// for element := data.Front(); element != nil; element = element.Next() {
	// 	fmt.Println(element.Value)
	// }

	for e := data.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
}
