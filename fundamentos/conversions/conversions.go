package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// Cuidado ... No codigo abaixo é convertido com base na tabela ascii
	fmt.Println("Teste" + string(120))
	// Int to string
	fmt.Println("Teste " + strconv.Itoa(123))

	// string to int
	num, _ := strconv.Atoi("2344")

	fmt.Println(num)
}
