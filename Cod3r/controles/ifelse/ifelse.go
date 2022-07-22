package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	aprovado := calcularNota(7.0)

	if aprovado {
		fmt.Println("Aprovado")
	} else {
		fmt.Println("Reprovado")
	}

	if i := gerarNumAleatorio(); i > 5 {
		fmt.Println("O numero é", i)
	}

}

func calcularNota(nota float64) bool {
	if nota >= 7 {
		return true
	} else {
		return false
	}
}

func notaConceito(nota float64) string {
	if nota >= 9 {
		return "A"
	} else if nota >= 7 {
		return "B"
	} else if nota >= 5 {
		return "C"
	} else if nota >= 3 {
		return "D"
	} else {
		return "E"
	}
}

func gerarNumAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}
