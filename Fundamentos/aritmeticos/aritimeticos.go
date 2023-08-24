package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma =>", a+b)
	fmt.Println("Subtração=>", a-b)
	fmt.Println("Divisão=>", a/b)
	fmt.Println("Multiplicação=>", a*b)
	fmt.Println("Módulo=>", a%b)

	//bitabit ou bitwide

	fmt.Println("E =>", a&b)   // 11 & 10 = 10   e
	fmt.Println("Ou =>", a|b)  // 11 | 10 = 01	 ou
	fmt.Println("Xor =>", a^b) // 11 ^ 10 = 01   ou exclusivo

	c := 3.0
	d := 2.0
	//outras operacoes usando math...

	fmt.Println("Maior =>", math.Max(float64(a), float64(b)))
	fmt.Println("Menor=>", math.Min(c, d))
	fmt.Println("Exponenciação =>", math.Pow(c, d))

}
