package main

import (
	"fmt"
	"time"
)

func comprar(t1, t2 bool) (bool, bool, bool) {

	if t1 {
		fmt.Println("eh")
	}

	return t1, t2, t1 != t2
}
func main() {
	fmt.Println(time.Unix(0, 0))

	type Pessoa struct {
		Nome string
		age  int16
	}
	p1 := Pessoa{"rodrigo", 0}
	fmt.Println(p1)

	fmt.Println(comprar(true, false))

	fmt.Println(time.Now().UnixNano())
}
