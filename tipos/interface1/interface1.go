package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

// Os métodos da interface imprimivel são implementados aqui
func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s e preço R$%.2f.", p.nome, p.preco)
}

// Utilizando a struct de forma polimorfica na função
func imprimir(p imprimivel) {
	fmt.Println(p.toString())
}

func main() {
	var p1 imprimivel = pessoa{"Felipe", "Henrique"}

	fmt.Println(p1.toString())
	imprimir(p1)

	p1 = produto{"Blusa", 17.90}
	fmt.Println(p1.toString())
	imprimir(p1)

	p2 := produto{"Calça jeans", 179.00}
	imprimir(p2)
}
