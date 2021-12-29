package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "Put Your Host")

	flag.Parse()

	fmt.Println(*host)
}
