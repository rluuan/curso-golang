package main

import (
	"fmt"
)

func main() {
	fmt.Print("mesma")
	fmt.Print(" linha.")

	x := 3.1415

	xs := fmt.Sprint(x)

	fmt.Println("O valor de x eh = " + xs)
	fmt.Println("O valor de x eh", x)

	fmt.Printf("O valor de x eh = %.2f", x)
	fmt.Printf("\nO valor de x eh = %s", xs)

	a, b, c, d := "rodrigo", false, 1.12342314, 't'

	fmt.Printf("\n%v %v %v %c", a, b, c, d)
}
