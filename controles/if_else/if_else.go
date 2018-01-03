package main

import "fmt"

func imprimeResultado(nota float64) {
	if nota > 7 {
		fmt.Println("aprovado")
	} else {
		fmt.Println("reprovado")
	}
}

func main() {
	imprimeResultado(7)
}
