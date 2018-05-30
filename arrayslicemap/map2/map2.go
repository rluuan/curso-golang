package main

import (
	"fmt"
)

func main() {
	funcESalarios := map[string]float64{
		"Jose":    497.2,
		"Rodrigo": 2423.44,
		"Rogerio": 434.43,
	}

	funcESalarios["Karina"] = 234.23

	for nome, sal := range funcESalarios {
		fmt.Println(nome, sal)
	}

	delete(funcESalarios, "Karina")
	for nome, sal := range funcESalarios {
		fmt.Println(nome, sal)
	}

}
