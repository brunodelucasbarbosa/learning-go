package main

import "fmt"

func main() {
	var frutas [2]string

	frutas[0] = "Banana"

	fmt.Println(frutas)

	frutas2 := []string{"Abacaxi", "Morango"}

	frutas2 = append(frutas2, "Abacaxi")
	fmt.Printf("frutas2: %v\n", frutas2)

	var numeros []int

	numeros = append(numeros, 1, 2, 3, 4, 5)

	fmt.Print(numeros)
}
