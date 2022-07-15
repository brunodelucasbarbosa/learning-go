package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hello, World!")
	var nome string = "Bruno"

	var idade float32 = 22
	var cidade string = "Salvador"

	nomeInferencia := "Bruno3"

	fmt.Println(nomeInferencia)

	if nome == "Bruno" {
		fmt.Println("Meu nome é ", nome)
		fmt.Println(`Meu nome é ${nome}`)
		fmt.Println(`Minha idade é `, idade)
		fmt.Println(`Minha cidade é `, cidade)
	}
	
	fmt.Println(reflect.TypeOf(nome))

	//function with return n * n
	funcionario := func(n int) int {
		return n * n
	}

	fmt.Println(funcionario(2))
}
