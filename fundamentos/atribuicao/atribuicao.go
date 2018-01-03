package main

import "fmt"

func main() {
	var b byte = 3
	fmt.Println(b)

	i := 3
	i += 3
	i -= 3
	i /= 2
	i *= 2
	i %= 2
	fmt.Println(b)
	// atribuicao de multiplos valores
	x, y := 3, 2
	fmt.Println("Primeiro era ", x, y)
	x, y = y, x
	fmt.Println("Depois ficou", x, y)
}
