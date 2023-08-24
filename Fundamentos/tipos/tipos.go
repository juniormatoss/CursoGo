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

	// representa um mapeamento da tbela Unicode(int32)
	var i2 rune = 'a' 
	fmt.Println("O uneé", reflect.TypeOf(i2))
	fmt.Println(i2)

	//números reais(float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))

	//boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	//String
	s1 := "Olá meu nome é júnior"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da String é", len(s1))

	//String com multiplas linhas
	s2 := `Olá
	meu
	nome é Leo`
	fmt.Print("O tamanho da string é", len(s2))

	//Char

	char := 'a'
	fmt.Println("O tipo de cha é", reflect.TypeOf(char))
	fmt.Println(char)

	
}
