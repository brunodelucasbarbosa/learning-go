package main

import "fmt"

func main() {

	nome,idade := retornaDoisValores()
	fmt.Println(nome, idade)

	exibirDados()
}

func exibirDados() {
	fmt.Println("Teste")
}

func retornaDoisValores() (string,int) {
	return "Bruno", 22
}