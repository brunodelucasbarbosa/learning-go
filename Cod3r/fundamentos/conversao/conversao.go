package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.5
	y := 2

	fmt.Println(int(x) / y)
	fmt.Println(x / float64(y))

	fmt.Println("comportamento estranho " + string(97))

	fmt.Println("int para string : " + strconv.Itoa(123))

	num, _ := strconv.Atoi("123")
	fmt.Println("string para int : ", num)
}
