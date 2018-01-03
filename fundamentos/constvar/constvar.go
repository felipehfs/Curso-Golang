package main

import (
	"fmt"
	m "math"
)

func main(){
	const Pi float64 = 3.1415
	var raio = 3.2
	area := Pi * m.Pow(raio, 2)
	fmt.Println(area)

}