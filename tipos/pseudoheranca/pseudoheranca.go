package main

import (
	"fmt"
)

type carro struct {
	nome  string
	marca string
}
type ferrari struct {
	carro // campos anonimos - meio que uma heranca implicita

	turboLigado bool
}

func main() {
	f := ferrari{}
	f.nome = "sei la1"
	f.marca = "sei la marca"
	f.turboLigado = true

	fmt.Println(f)
}
