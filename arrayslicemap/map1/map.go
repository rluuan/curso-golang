package main

import (
	"fmt"
)

func main() {
	// var aprovados map[int]string

	//  mapas devem ser inicializados

	aprovados := make(map[int]string)

	aprovados[1234567890] = "Maria"
	aprovados[1234567891] = "Carlos"
	aprovados[1234567892] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Println(nome, " cpf = ", cpf)
	}

	fmt.Println(aprovados[1234567892])
	delete(aprovados, 1234567892) // chave do map para excluir
	fmt.Println(aprovados)
}
