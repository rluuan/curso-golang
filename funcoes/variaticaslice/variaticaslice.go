package main

import (
	"fmt"
)

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de aprovados:")

	for i, valor := range aprovados {
		fmt.Printf("%d) %s\n", i+1, valor)
	}
}

func main() {
	aprovados := []string{"Rodrigo", "Luan", "Elenita", "Genivado"}
	imprimirAprovados(aprovados...)

}
