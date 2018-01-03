package main

import "fmt"

func main() {
	// Os mapas precisam ser inicializados
	aprovados := make(map[int]string)

	aprovados[1234040222] = "Anna"
	aprovados[4443442234] = "Washington"
	aprovados[6073443322] = "Cleber"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}
	// Remove um item do mapa
	delete(aprovados, 1234040222)
}
