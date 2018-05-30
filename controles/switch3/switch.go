package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "integer"
	case float32, float64:
		return "float"
	case string:
		return "string"
	case func():
		return "function"
	default:
		return "I dont know"
	}
}

func main() {
	fmt.Println(tipo(98237423))
	fmt.Println(tipo(1.423))
	fmt.Println(tipo("rodrigo"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))

}
