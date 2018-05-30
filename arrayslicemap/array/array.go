package main

import (
	"fmt"
	"strings"
)

func main() {

	var notas [3]float64
	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 7.8, 2, 3.3

	s := strings.Split("rodrigo, luan, santos, lima", "")
	fmt.Println(s)

	fmt.Println(len(notas))

}
