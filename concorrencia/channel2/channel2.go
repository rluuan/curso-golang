package main

import (
	"fmt"
	"time"
)

// channel eh um tipo
// Channel (canal) - eh a forma de comunicacao entre garoutines

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base

	time.Sleep(2 * time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base

}
func main() {
	c := make(chan int)

	go doisTresQuatroVezes(2, c)

	a, b := <-c, <-c // recebendo dados do canal

	fmt.Println(a, b)

	fmt.Println(<-c)
}
