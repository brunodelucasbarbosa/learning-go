package main

import "fmt"

func main() {
	var nome string
	var idade int
	fmt.Println("Digite seu nome: ")
	fmt.Scanf("%s", &nome)
	fmt.Println("Digite sua idade: ")
	fmt.Scan(&idade)

	fmt.Println("O nome digitado foi: ", nome)
	fmt.Println("A idade digitada foi: ", idade)
}