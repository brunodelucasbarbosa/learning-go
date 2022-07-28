package main

import "fmt"

func main() {
	p1 := Ponto{x: 2, y: 2}
	p2 := Ponto{x: 2, y: 4}

	fmt.Println(catetos(p1, p2))

	fmt.Println(Distancia(p1, p2))
}
