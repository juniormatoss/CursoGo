package main

import "fmt"

func main() {
	var b byte = 3
	fmt.Println(b)

	i := 3 // inrefência de tipo - ja inicializa
	i += 3 // esta acrescentando +3 na váriavel i
	i -= 3 // está retirando -3 da várivel i
	i /= 2 // i = i/2
	i *= 2 // i = i*2
	i %= 2 // i = i%2

	fmt.Println(i)

	x, y := 1, 2
	fmt.Println(x, y)

	x, y = y, x
	fmt.Println(x, y)
}
