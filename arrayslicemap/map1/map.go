package main

import "fmt"

func main() {

	// var aprovados map[int]string
	// mapas devem ser inicializados

	aprovados := make(map[int]string)

	aprovados[12345678978] = "Matia"
	aprovados[98765432100] = "Pedro"
	aprovados[46574746874] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[46574746874])
	delete(aprovados, 46574746874)
	fmt.Println(aprovados[46574746874])
}
