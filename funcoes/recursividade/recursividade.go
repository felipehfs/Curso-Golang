package main

import "fmt"

func fatorial(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		return fatorial(n-1) * n
	}
}

func main() {
	result := fatorial(5)
	fmt.Println(result)
}
