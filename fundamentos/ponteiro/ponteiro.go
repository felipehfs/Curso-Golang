package main

import "fmt"

func main() {
	i := 1
	// Go não tem aritmética de ponteiros
	var p *int
	p = &i // Pegando valor da referência
	*p++
	i++
	fmt.Println(p, *p, i)
}
