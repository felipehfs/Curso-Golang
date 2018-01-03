package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// Numeros inteiros
	fmt.Println(1, 1000, 2)
	fmt.Println("INSPECT: ", reflect.TypeOf(38990))

	// sem sinal (só positivos)... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O valor do byte é", reflect.TypeOf(b))

	i1 := math.MaxInt64
	fmt.Println("O valor maximo do int64", i1)
	fmt.Println("O tipo do i1 é ", reflect.TypeOf(i1))

	b1 := false
	fmt.Println("O valor de b1 é", reflect.TypeOf(b1))

	// String com múltiplas linhas
	str1 := `
		Olá meu nome é
		Felipe
	`
	fmt.Println(str1)

	// Char não existe em Go, na verdade é um valor int32
	char := 'b'
	fmt.Println("O real valor do char é", reflect.TypeOf(char))
}
