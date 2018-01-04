package main

import "fmt"

func troca(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
	primeiro = p1
	return // Retorno limpo
}

func main() {
	primeiro, segundo := troca(20, 43)
	fmt.Println(primeiro, segundo)
}
