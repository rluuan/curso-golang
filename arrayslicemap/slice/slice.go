package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3}
	s1 := [...]int{4, 5, 6}
	fmt.Println(a1, s1, reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := a1[:2]
	fmt.Println(a1, a2)
}
