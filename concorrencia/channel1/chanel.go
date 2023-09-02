package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 1 // enviando dados para o canal (escrita)
	<-ch    // recebendo dasdos do canal(leitura)

	ch <- 2
	fmt.Println(<-ch)

}
