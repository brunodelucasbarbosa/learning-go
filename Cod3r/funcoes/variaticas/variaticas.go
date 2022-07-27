package main

import "fmt"

func media(notas ...float64) float64 {
	total := 0.0

	for _, nota := range notas {
		total += nota
	}

	return total / float64(len(notas))
}

func main() {
	fmt.Println(media(7.7, 8.7, 9.7))
	fmt.Println(media(7.7, 8.7, 9.7, 10.7))
	fmt.Println(media(7.7, 8.7, 9.7, 10.7, 11.7))
}
