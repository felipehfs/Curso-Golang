package main

import (
	"fmt"
)

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// Método: função com receiver (receptor)
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var produto1 produto
	produto1 = produto{
		nome:     "Tesoura",
		preco:    4.30,
		desconto: 0.05,
	}

	fmt.Println(produto1, produto1.precoComDesconto())
	produto2 := produto{"Notebook", 1890.00, 0.03}
	fmt.Println(produto2, produto2.precoComDesconto())
}
