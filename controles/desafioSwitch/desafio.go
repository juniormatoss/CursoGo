package main

import "fmt"

func notaParaConceito(n float64) string {
	switch {
	case n >= 8 && n < 9:
		return "B"
	case n >= 5 && n < 8:
		return "C"
	case n >= 3 && n < 5:
		return "B"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(notaParaConceito(7))
}
