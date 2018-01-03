package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // compilador conta!

	for index, numero := range numeros {
		fmt.Printf("%d) %d\n", index+1, numero)
	}
}
