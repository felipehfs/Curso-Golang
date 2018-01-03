package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José joão":      12224.53,
		"Gabriela Silva": 15533.78,
		"Pedro Junior":   1200.06,
	}
	funcsESalarios["Rafael Filho"] = 2000.04

	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
