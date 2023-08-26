package main

import "fmt"

func main() {

	var notas [3]float64

	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 7, 4, 10

	fmt.Println(notas)

	total := 0.0
	fmt.Println("Primeiro Total =" , total)

	for i:= 0; i <len(notas); i++{
		total = total + notas[i]
	}
	fmt.Println("segundo total= ", total/3)
}
