package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2

	area := PI * math.Pow(raio, 2)
	fmt.Println("a área da circuferencia é:", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)
	fmt.Println(a, b, c, d)
 
	g, h, i := 10, true, " epa!"
	fmt.Print(g, h, i)

}
