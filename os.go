package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Argument : ", args)

	name, error := os.Hostname()

	if error == nil {
		fmt.Println("HostName : ", name)
	} else {
		fmt.Println("Error : ", error)
	}
}
