package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	switch {
	case now.Hour() > 12:
		fmt.Println("Boa tarde")
	case now.Hour() >= 18:
		fmt.Printf("Boa noite")
	default:
		fmt.Println("Bom dia")
	}
}
