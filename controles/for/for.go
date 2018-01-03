package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for i <= 10 { // Não tem while no Go
		fmt.Println(i)
		i++
	}

	for j := 0; j <= 20; j++ {
		fmt.Printf("%d\n", j)
	}

	for {
		fmt.Println("Laço infinito...")
		time.Sleep(time.Second)
	}
}
