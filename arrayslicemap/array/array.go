package main

import "fmt"

func main() {
	// homogênea (mesmo tipo) e estática (fixo)
	var notas [3]float64
	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 7.9, 4.3, 9.1
	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}
	total /= float64(len(notas))
	fmt.Println(total)
}
