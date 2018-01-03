package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	return random.Intn(10)
}

func main() {
	if i := numeroAleatorio(); i > 5 { // i visível apenas no bloco
		fmt.Println("Ganhou!")
	} else {
		fmt.Println("Perdeu!")
	}
}
