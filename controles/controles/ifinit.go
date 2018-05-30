package main

import (
	"fmt"
	"math/rand"
	"time"
)

func one() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}
func p(t string) {
	fmt.Println(t)
}
func main() {
	if i := one(); i > 5 {
		p("Ganhou")
	} else {
		p("Perdeu")
	}

}
