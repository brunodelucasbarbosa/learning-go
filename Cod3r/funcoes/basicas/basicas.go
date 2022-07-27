package main

import "fmt"

func imprimirOi() {
	fmt.Println("Oi")
}

func retornarOi() string {
	return "Oi"
}

func retornarOQuadrado(x int) int {
	return x * x
}

func retornarOiTudoBem() (string, string) {
	return "Oi", "Tudo bem?"
}

var soma = func(a int, b int) (soma int) {
	soma = a + b
	return
}

func main() {
	imprimirOi()
	fmt.Println(retornarOi())
	fmt.Println(retornarOQuadrado(2))
	fmt.Println(retornarOiTudoBem())

	oi, tudoBem := retornarOiTudoBem()

	fmt.Print(oi, " ", tudoBem)
}
