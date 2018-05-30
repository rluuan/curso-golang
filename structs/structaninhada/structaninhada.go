package main

import (
	"fmt"
)

type item struct {
	produtoID int
	qtd       int
	preco     float64
}
type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtd)
	}
	return total
}
func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			item{produtoID: 1, qtd: 2, preco: 20.4},
			item{produtoID: 2, qtd: 1, preco: 30.4},
			item{produtoID: 3, qtd: 1, preco: 40.4},
		},
	}
	fmt.Println(pedido.valorTotal())
}
