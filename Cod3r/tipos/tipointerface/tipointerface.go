package main

import "fmt"

type curso struct {
	nome string
}

func main() {
	var coisa interface{}

	coisa = 3
	fmt.Println(coisa)

	type dinamico interface{}

	var coisa2 dinamico = "Opa"
	fmt.Println(coisa2)

	coisa2 = curso{nome: "Go"}
	fmt.Println(coisa2)
}
