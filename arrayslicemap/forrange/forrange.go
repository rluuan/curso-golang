package main

import (
	"fmt"
)

func main() {
	// numeros := [...]int{10, 20, 30, 40, 50}
	numeros := "rodrigo"
	for i, value := range numeros {

		fmt.Println(i, string(value))
	}
}
