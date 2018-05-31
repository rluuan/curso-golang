package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	// struct para json

	p1 := produto{1, "notebook", 1899.90, []string{"promocao", "eletronico"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	// json para struct
	var p2 produto
	jsonToString := `{"id":1,"nome":"notebook","preco":1899.9,"tags":["promocao","eletronico"]}`
	json.Unmarshal([]byte(jsonToString), &p2)
	fmt.Println(p2)
}
