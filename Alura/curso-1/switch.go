package main

import "fmt"

func main() {

	var idade int
	fmt.Print("Digite sua idade: ")
	fmt.Scan(&idade)

	switch idade {
		case 18:
			fmt.Println("Você tem 18 anos")
		case 20:
			fmt.Println("Você tem 20 anos")
		
		default:
			fmt.Println("Você não tem 18 nem 20 anos")
	}
}