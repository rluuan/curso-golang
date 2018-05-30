package main

import (
	"fmt"
)

func obterValor(numero int) int {
	defer fmt.Println("fim")

	if numero > 1000 {
		fmt.Println("maior")
		return numero
	} else {
		fmt.Println("menor ou igual")
		return 1000
	}
}

func main() {
	fmt.Println(obterValor(500))
}
