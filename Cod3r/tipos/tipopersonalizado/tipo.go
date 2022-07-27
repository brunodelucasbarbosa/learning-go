package main

import "fmt"

type nota float64

func (n nota) entre(inicio float64, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
}

func notaParaConceito(n nota) string {
	if n.entre(9.0, 10.0) {
		return "A"
	}
	if n.entre(7.9, 8.9) {
		return "B"
	}
	if n.entre(5.7, 6.9) {
		return "C"
	}
	if n.entre(3.5, 4.9) {
		return "D"
	}

	return "E"
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(2.9))
}
