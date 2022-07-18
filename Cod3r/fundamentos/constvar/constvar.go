package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	raio := 3.2

	area := PI * math.Pow(raio, 2)

	fmt.Printf("A área da circunferência é %.2f", area)
}
