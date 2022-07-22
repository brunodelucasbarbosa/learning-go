package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	num := 2

	switch num {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("Um") // Como o switch não tem break, o case 1 será executado
	case 2:
		fmt.Println("Dois")
	default:
		fmt.Println("Não é nenhum dos casos")
	}

	switch {
	case num < 5:
		fmt.Println("Menor que 5")
	case num > 5:
		fmt.Println("Maior que 5")
	}

}

func gerarNumAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}
