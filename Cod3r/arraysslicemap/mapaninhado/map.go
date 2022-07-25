package main

import "fmt"

func main() {
	funcsESalarios := map[string]map[string]float64{
		"G": {
			"Gabriel João":   11325.45,
			"Gabriela Silva": 15456.78,
			"Giulio Junior":  1200.00,
		},
		"P": {
			"Pedro Silva":   11325.45,
			"Paulo Silva":   15456.78,
			"Pamela Junior": 1200.00,
		},
	}

	fmt.Println(funcsESalarios)

}
