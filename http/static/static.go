package main 

import (
	"net/http"
	"log"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
