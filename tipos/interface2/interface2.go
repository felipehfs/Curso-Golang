package main

/*
 Exercício de uma pegadinha muito interessante em Go.
 Quando utilizamos o polimorfismo em um método que tem efeito colateral,
 é necessário indicar o endereço da struct.
 Além disso, é sempre uma boa prática implementar interfaces que não tenham
 um efeito colateral
*/
type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	nome  string
	turbo bool
}

func (f *ferrari) ligarTurbo() {
	f.turbo = true
}

func main() {
	var car1 esportivo = &ferrari{"F39", false}
	car1.ligarTurbo()
}
