package main

import "fmt"

func main() {
	a := 10

	var p *int = nil

	p = &a
	a += 10
	fmt.Println(*p) // obter o valor do ponteiro p
	fmt.Println(p) // obter o endereço do ponteiro p 
	fmt.Println(a)

}
