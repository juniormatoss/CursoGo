package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	//sem sinal(só positivos)  uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	// com sinal int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo o int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(12))


	var i2 rune = 'a' // representa um mapeamento da tbela Unicode(int32)
	fmt.Println("O uneé", reflect.TypeOf(i2))
	fmt.Println(i2)



	//números reais(float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))

	//boolean
	 bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

}
