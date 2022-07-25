package main

import (
	"fmt"
	"math/rand"
	"time"
)

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case bool:
		return "booleano"
	case func():
		return "função"
	default:
		return "não sei"
	}
}

func main() {
	fmt.Println(tipo(1))
	fmt.Println(tipo(1.1))
	fmt.Println(tipo("string"))
	fmt.Println(tipo(true))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(rand.New(rand.NewSource(time.Now().UnixNano()))))
}