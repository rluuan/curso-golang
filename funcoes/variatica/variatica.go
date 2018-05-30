package main

import (
	"fmt"
)

func variatica(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	return total / float64(len(numeros))
}

func main() {
	fmt.Println(variatica(5.5, 5.5, 2.2))

}
