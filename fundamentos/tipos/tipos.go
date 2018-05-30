package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro eh:", reflect.TypeOf(1000))

	var b byte = 255

	fmt.Println(reflect.TypeOf(b))

	i1 := math.MaxInt64

	fmt.Println(i1, reflect.TypeOf(i1))

	var i2 rune = 'a'
	fmt.Println(reflect.TypeOf(i2))
	fmt.Println(i2)

	s1 := "rodrigo luan"
	fmt.Println(len(s1))

	s2 := `rodrigo
	luan
	teste`
	fmt.Println(s2)
}
