package main

import "fmt"

func imprimitResultado(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado", nota)

	} else {
		fmt.Println("Reprovado", nota)
	}
}
func main(){
	imprimitResultado(9.0)
}
