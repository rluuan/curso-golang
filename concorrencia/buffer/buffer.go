package main

import (
	"fmt"
)

func rotina(c chan int) {
	c <- 1
	c <- 2
	c <- 3
	fmt.Println("Executou")
	c <- 4
	c <- 5
	c <- 1
}

func main() {
	c := make(chan int, 3)
	go rotina(c)
	fmt.Println(<-c)
}
