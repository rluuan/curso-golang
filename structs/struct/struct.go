package main

import (
	"fmt"
	"strings"
)

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var p produto
	p = produto{
		"Rodrigo",
		1234.44,
		0.10,
	}
	fmt.Println(strings.Split(p.nome, ""), p.precoComDesconto())
}
