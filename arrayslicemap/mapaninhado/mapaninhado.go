package main

import (
	"fmt"
)

func main() {

	funcPorLetras := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 4324.34,
			"Maria Silva":    231.34,
		},
		"E": {
			"Pedro Silva":  321.34,
			"Marcos Silva": 1324.34,
		},
		"I": {
			"Hugo Silva":  132.34,
			"Kaick Silva": 2344.34,
		},
	}
	delete(funcPorLetras, "G")

	for letra, funcs := range funcPorLetras {
		fmt.Println(letra, funcs)
	}
}
