package database

import "fmt"

var connection string

// Init akan selalu dieksekusi ketika package diimport
func init() {
	fmt.Println("Init Dipanggil")
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
