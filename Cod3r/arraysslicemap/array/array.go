package main

import (
	"fmt"
)

func main() {
	var notas [3]float64

	notas[0] = 7.0
	notas[1] = 8.0
	notas[2] = 9.0

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))

	fmt.Println(notas)
	fmt.Println(media)
}