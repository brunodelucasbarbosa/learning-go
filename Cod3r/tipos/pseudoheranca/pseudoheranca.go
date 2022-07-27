package main

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro   // campo anonimo
	turboOn bool
}

func main() {
	f := ferrari{turboOn: true}
	f.nome = "F40"
	f.velocidadeAtual = 0

	fmt.Println(f)
}
