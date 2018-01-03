package main

import "fmt"

func somar(x int, y int) int {
	return x + y
}
func imprimir(x int) {
	fmt.Println(x)
}

func main() {
	var total = somar(3, 6)
	imprimir(total)
}
