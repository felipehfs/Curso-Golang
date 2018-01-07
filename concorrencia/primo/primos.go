package main

import (
	"fmt"
)

// isPrime verfica se o número é primo
func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n % i == 0{
			return false		
		}
	}
	return true
}

// primos retorna um conjunto de números primos pelo channel
func primos(n int, c chan int) {
	inicio := 2
	for i := 0; i < n; i++ {
		for primo := inicio; ; primo++ {
			if isPrime(primo) {
				c <- primo
				inicio = primo + 1
				break
			}
		}
	}
	close(c)
}

func main() {
	c := make(chan int, 30)
	go primos(60, c)
	for primo := range c {
		fmt.Printf("%d ", primo)
	}
	fmt.Println("Fim!")
}
