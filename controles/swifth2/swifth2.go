package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()
	switch { // ele vai imprimir o primeiro que for verdade
	case t.Hour() < 12:
		fmt.Println("Bom dia!")

	case t.Hour() < 18:
		fmt.Println("Boa tarde")

	default:
		fmt.Println("Boa")

	}
}
