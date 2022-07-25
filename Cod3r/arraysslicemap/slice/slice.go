package main

import (
	"fmt"
)

func main() {
	a1 := [3]int{1, 2, 3} // array de tamanho 3
	a2 := []int{1, 2, 3}  // slice de tamanho 3

	fmt.Println(a1)
	fmt.Println(a2)

	a3 := [5]int{1, 2, 3, 4, 5}

	s3 := a3[0:3] // slice de tamanho 2
	fmt.Println(s3)
}
