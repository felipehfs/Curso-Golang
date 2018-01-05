package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
}

type ferrari struct{}

func (f ferrari) fazerBaliza() {
	fmt.Println("Fazendo baliza")
}

func (f ferrari) ligarTurbo() {
	fmt.Println("Ligando o turbo")
}

func main() {
	var car esportivoLuxuoso = ferrari{}
	car.fazerBaliza()
	car.ligarTurbo()
}
