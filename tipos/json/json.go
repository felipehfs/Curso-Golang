package main

import (
	"encoding/json"
	"fmt"
)

/*
	Convertendo Json para struct e vice-versa...
*/

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	p1 := produto{1, "Camiseta", 30.00, []string{"Vestimenta", "Verão"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	var p2 produto
	jsonString := `{"id": 1, "nome":"Ps4", "preco": 1330.00, "tags":["games", "console"] }`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2)
}
