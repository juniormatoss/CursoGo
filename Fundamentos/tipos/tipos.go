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

	
}
