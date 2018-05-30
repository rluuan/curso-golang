package main

import (
	"fmt"
)

func trocar(p1, p2 int) (segundo, primeiro int) {
	segundo = p2
	primeiro = p1
	return
}

func main() {
	x, y := trocar(3, 2)
	fmt.Println(x, y)

	first, second := trocar(7, 1)
	fmt.Println(first, second)
}
