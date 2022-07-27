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
	// struct to json

	p1 := produto{
		1,
		"Notebook",
		1989.90,
		[]string{
			"eletronico",
			"gamer",
		},
	}

	p1ToJson, _ := json.Marshal(p1)

	fmt.Println(p1)
	fmt.Println(string(p1ToJson))

	p2 := produto{}
	jsonToString := `{"id":1,"nome":"Notebook","preco":1989.90,"tags":["eletronico","gamer"]}`
	json.Unmarshal([]byte(jsonToString), &p2)

	fmt.Println(p2.Tags[0])
}
