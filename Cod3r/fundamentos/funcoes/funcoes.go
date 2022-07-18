package main

import "fmt"

func main() {
	fmt.Println(somar(2, 3))
}

func somar(a int, b int) (string, int) {
	return "O resultado é:", a + b 
}
