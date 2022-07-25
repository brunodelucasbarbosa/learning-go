package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José João":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.00,
	}
	delete(funcsESalarios, "José João")
	fmt.Println(funcsESalarios)

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
