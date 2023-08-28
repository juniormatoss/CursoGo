package main

import "fmt"

// primeiro float64 indica que a entrada precisa ser em float
// o segundo indica que a saida também sera em float64
func media(numeros ...float64) float64 {
	total := 0.0

	for _, num := range numeros {
		total += num
	}
	return total / float64(len(numeros))
}

func main() {
	fmt.Printf("Média: %.2f", media(7.7, 8.1, 5.9))
}
