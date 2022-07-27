package main

import "fmt"

func multiplicar(a int, b int) int {
	return a * b
}

func exec(funcao func(int, int) int, p1 int, p2 int) int {
	return funcao(p1, p2)
}

func main() {
	fmt.Println(exec(multiplicar, 2, 3))
}
