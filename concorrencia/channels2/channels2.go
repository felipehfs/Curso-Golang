package main

import (
	"fmt"
	"time"
)

// Channels (canal) - é a forma de comunicação entre goroutines
// channel é um tipo

func doistTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // enviando dados para o canal
	// Enquanto a informação não é lida o programa trava nessa linha
	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}
func main() {
	c := make(chan int)
	go doistTresQuatroVezes(2, c)

	a, b := <-c, <-c //recebendo dados do canal
	fmt.Println(a, b)

	fmt.Println(<-c)
}
