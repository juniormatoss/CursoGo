package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)


	//cuidado obs: ele vai printar o valor na tabela ask e mostra
	fmt.Println("Teste " + string(97))

	//int para string
	fmt.Println("Teste", strconv.Itoa(123))

	//String para int obs: o _ serve para ignorar pois ele retorna dois resultados e o _ ignora o segundo
	num, _ := strconv.Atoi("123")
	fmt.Println(num -122)

	b, _ := strconv.ParseBool("true")
	if b{
		fmt.Println("Verdadeiro")
	}
}
