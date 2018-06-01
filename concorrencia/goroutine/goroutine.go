package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtd int) {
	for i := 0; i < qtd; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteracao %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// fale("Maria", "Porque voce nao fala comigo", 3)
	// fale("Joao", "So posso falar depois de voce", 1)

	// go fale("Maria", "Ei...", 500)
	// go fale("Joao", "Opa...", 500)

	go fale("Maria", "entendi", 10)
	fale("Joao", "parabens!", 5)
}
