package main

import "fmt"

func main() {

	// Declare Array
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	// Declare Slice
	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// Mengubah Slice Akan mempengaruhi array, begitupun sebaliknya
	months[5] = "Diubah"
	fmt.Println(slice1)

	slice1[1] = "Juni"
	fmt.Println(slice1)

	var slice2 = months[10:]
	fmt.Println(slice2)

	// append akan membuat array baru, yg lama tidak akan berubah
	var slice3 = append(slice2, "Bulan Baru")
	fmt.Println(slice3)

	slice3[1] = "Bukan Desember"

	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(months)

	// Declare slice langsung tanpa membuat array terlebih dahulu
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Ravi"
	newSlice[1] = "Mukti"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// Copy Slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	// hati hati declarasi array tanpa length, karena nanti akan menjadi tipe data slice
	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
