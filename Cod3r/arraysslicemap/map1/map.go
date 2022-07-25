package main

import (
	"fmt"
)

func main() {
	aprovados := make(map[int]string)

	aprovados[12345] = "Maria"
	aprovados[67890] = "Pedro"
	aprovados[54321] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Println("CPF: ", cpf, " - Nome: ", nome)
	}
	delete(aprovados, 12345)
	fmt.Println(aprovados)
}
