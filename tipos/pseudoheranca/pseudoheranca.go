package main

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro       // campos anonimos
	turboLigado bool
}

func main() {
	f := ferrari{}
	f.nome = "F50"
	f.velocidadeAtual = 0
	f.turboLigado = true

	fmt.Printf("A ferrari %s está com o turbo ligado? %v", f.nome, f.turboLigado)
}
