package main

import "fmt"

func obterValorAprovado(numero int) int {
	defer fmt.Println("Saindo do metodo...")
	if numero > 5 {
		fmt.Println("Valor maior que 5")
		return numero
	}
	fmt.Println("Valor retornado:", numero)
	return numero
}

func main() {
	fmt.Println(obterValorAprovado(6))
	fmt.Println(obterValorAprovado(4))
}
