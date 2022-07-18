package main

import "fmt"

func obterResultado(nota float64) string {
	if nota >= 7 {
		return "Aprovado"
	} else if nota >= 5 {
		return "Recuperação"
	} else {
		return "Reprovado"
	}
}

func main() {
	fmt.Println(obterResultado(6.9))
	fmt.Println(obterResultado(4.9))
	fmt.Println(obterResultado(2.9))
}
