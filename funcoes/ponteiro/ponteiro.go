package main

import "fmt"

func inc1(n int) {
	n++
}

// resisão: um ponteiro é representado por um *

func inc2(n *int) {
	// revisão: * é usado para desreferencias, ou seja
	// ter acesso ao valor no qual o ponteiro aponta
	*n++

}
func main() {
	n := 1

	inc1(n) // por valor
	fmt.Println(n)

	// Revisão & é uAdo para obeter o endereço da variável
	inc2(&n) // por referencia
	fmt.Println(n)
}
