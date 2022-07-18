package main

import (
	"fmt"
	"math"
)

func main() {
	a := 10
	b := 20

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a / b)

	fmt.Println(math.Max(float64(a), float64(b)))
	fmt.Println(math.Min(float64(a), float64(b)))
}
