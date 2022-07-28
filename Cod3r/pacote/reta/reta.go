package main

import "math"

type Ponto struct {
	// Ponto x, y float64
	x, y float64
}

func catetos(a Ponto, b Ponto) (cx float64, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)

	return cx, cy
}

func Distancia(a, b Ponto) float64 {
	// Distancia linear entre dois pontos
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
