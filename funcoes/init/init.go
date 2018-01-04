package main

import (
	"fmt"
)

// Por padrao a funcao init é sempre executado primeiro no package main
// antes da função main propriamente dita
func init() {
	fmt.Println("Initialing...")
	fmt.Println("Generating logging...")
}

func main() {
	fmt.Println("Main function.../")
}
