package main

import "fmt"

func funcao1() {
	fmt.Println("Função função1!")
}

func funcao4(p1, p2 string) string {
	return fmt.Sprintf("P1: %s P2: %s", p1, p2)
}

// Múltiplos retornos
func funcao5() (string, string) {
	return "Retorno 1", "Retorno 2"
}

func main() {
	funcao1()
	retorno1, retorno2 := funcao5()
	fmt.Println(retorno1, retorno2)
}
